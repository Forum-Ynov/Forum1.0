<?php

class MessagesController
{
    public function __construct(private MessagesGateway $gateway)
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

    public function processRequestByTopics(string $method, ?string $pseudo): void
    {
        if ($pseudo) {

            $this->processRessourceTopics($method, $pseudo);

        }
    }

    private function processRessourceRequest(string $method, string $id): void
    {
        $message = $this->gateway->get($id);

        if (!$message) {
            http_response_code(404);
            echo json_encode(["message" => "Message not found"], JSON_PRETTY_PRINT);
            return;
        }

        switch ($method) {
            case "GET":
                echo json_encode($message, JSON_PRETTY_PRINT);
                break;

            case "PATCH":
                $data = (array) json_decode(file_get_contents("php://input"), true);

                $rows = $this->gateway->update($message, $data);

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

    private function processRessourceTopics(string $method, string $topics): void
    {
        $message = $this->gateway->getByTopics($topics);

        if (!$message) {
            http_response_code(404);
            echo json_encode(["message" => "Message not found"], JSON_PRETTY_PRINT);
            return;
        }

        switch ($method) {
            case "GET":
                echo json_encode($message, JSON_PRETTY_PRINT);
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

                $id = $this->gateway->create($data);

                http_response_code(201);

                $this->processRessourceRequest("GET", $id);

                break;

            default:
                http_response_code(405);
                header("Allow: GET, POST");
        }
    }
   
}
