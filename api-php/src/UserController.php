<?php

class UserController
{
    public function __construct(private UserGateway $gateway)
    {
    }

    public function processRequest(string $method, ?string $id): void
    {
        if ($id) {

            $this->processRessourceRequest($method, $id);

        } else {

            $this->processCollectionRequest($method);

        }
    }

    public function processRequestByPseudo(string $method, ?string $pseudo): void
    {
        if ($pseudo) {

            $this->processRessourcePseudo($method, $pseudo);

        }
    }

    public function processRequestByEmail(string $method, ?string $email): void
    {
        if ($email) {

            $this->processRessourceEmail($method, $email);

        }
    }

    public function processRequestLogin(string $method): void
    {

        $this->processLogin($method);

    }

    private function processRessourceRequest(string $method, string $id): void
    {
        $user = $this->gateway->get($id);

        if (!$user) {
            http_response_code(404);
            echo json_encode(["message" => "User not found"], JSON_PRETTY_PRINT);
            return;
        }

        switch ($method) {
            case "GET":
                echo json_encode($user, JSON_PRETTY_PRINT);
                break;

            case "PATCH":
                $data = (array) json_decode(file_get_contents("php://input"), true);

                $errors = $this->getValidationErrors($data, false);

                if (!empty($errors)) {
                    http_response_code(422);
                    echo json_encode(["errors" => $errors], JSON_PRETTY_PRINT);
                    break;
                }

                $rows = $this->gateway->update($user, $data);

                $this->processRessourceRequest("GET", $id);

                break;

            case "DELETE":
                $rows = $this->gateway->delete($id);

                $this->processCollectionRequest("GET");
                break;

            default:
                http_response_code(405);
                header("Allow: GET, PATCH, DELETE");
        }

    }

    private function processRessourcePseudo(string $method, string $pseudo): void
    {
        $user = $this->gateway->getByPseudo($pseudo);

        if (!$user) {
            http_response_code(404);
            echo json_encode(["message" => "User not found"], JSON_PRETTY_PRINT);
            return;
        }

        switch ($method) {
            case "GET":
                echo json_encode($user, JSON_PRETTY_PRINT);
                break;

            default:
                http_response_code(405);
                header("Allow: GET");
        }

    }

    private function processRessourceEmail(string $method, string $email): void
    {
        $user = $this->gateway->getByEmail($email);

        if (!$user) {
            http_response_code(404);
            echo json_encode(["message" => "User not found"], JSON_PRETTY_PRINT);
            return;
        }

        switch ($method) {
            case "GET":
                echo json_encode($user, JSON_PRETTY_PRINT);
                break;

            default:
                http_response_code(405);
                header("Allow: GET");
        }

    }

    private function processCollectionRequest(string $method): void
    {
        switch ($method) {
            case "GET":
                echo json_encode($this->gateway->getAll(), JSON_PRETTY_PRINT);
                break;
            case "POST":
                $data = (array) json_decode(file_get_contents("php://input"), true);

                $errors = $this->getValidationErrors($data);

                if (!empty($errors)) {
                    http_response_code(422);
                    echo json_encode(["errors" => $errors], JSON_PRETTY_PRINT);
                    break;
                }

                $user = $this->gateway->getByPseudo($data["pseudo"]);

                if ($user) {
                    http_response_code(404);
                    echo json_encode(["message" => "pseudo already used "], JSON_PRETTY_PRINT);
                    break;
                }

                $user = $this->gateway->getByEmail($data["email"]);

                if ($user) {
                    http_response_code(404);
                    echo json_encode(["message" => "email already used "], JSON_PRETTY_PRINT);
                    break;
                }

                $id = $this->gateway->create($data);

                http_response_code(201);

                $this->processRessourceRequest("GET", $id);

                break;

            default:
                http_response_code(405);
                header("Allow: GET, POST");
        }
    }

    private function processLogin(string $method): void
    {
        switch ($method) {
            case "POST":
                $data = (array) json_decode(file_get_contents("php://input"), true);

                $errors = $this->getValidationErrors($data);

                if (!empty($errors)) {
                    http_response_code(422);
                    echo json_encode(["errors" => $errors], JSON_PRETTY_PRINT);
                    break;
                }

                $user = $this->gateway->getByPseudo($data["pseudo"]);

                if (!$user) {
                    http_response_code(404);
                    echo json_encode(["message" => "user not found "], JSON_PRETTY_PRINT);
                    break;
                }

                if (!password_verify($data["passwd"], $user["passwd"])) {
                    http_response_code(404);
                    echo json_encode(["message" => "password incorect"], JSON_PRETTY_PRINT);
                    break;
                }

                echo json_encode($user, JSON_PRETTY_PRINT);
                break;

            default:
                http_response_code(405);
                header("Allow: POST");
        }
    }

    private function getValidationErrors(array $data, bool $is_new = true): array
    {
        $errors = [];

        if ($is_new && empty($data["pseudo"])) {
            $errors[] = "pseudo is required";
        }

        return $errors;
    }
}
