<?php

class TopicsController
{
    public function __construct(private TopicsGateway $gateway)
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

    public function processRequestByTags(string $method, ?string $tags): void
    {
        if ($tags) {

            $this->processRessourceTags($method, $tags);

        }
    }

    private function processRessourceRequest(string $method, string $id): void
    {
        $topics = $this->gateway->get($id);

        if (!$topics) {
            http_response_code(404);
            echo json_encode(["message" => "Topics not found"], JSON_PRETTY_PRINT);
            return;
        }

        switch ($method) {
            case "GET":
                echo json_encode($topics, JSON_PRETTY_PRINT);
                break;

            default:
                http_response_code(405);
                header("Allow: GET");
        }

    }

    private function processRessourceTags(string $method, string $tags): void
    {
        $topics = $this->gateway->getByTags($tags);

        if (!$topics) {
            http_response_code(404);
            echo json_encode(["message" => "Topics not found"], JSON_PRETTY_PRINT);
            return;
        }

        switch ($method) {
            case "GET":
                echo json_encode($topics, JSON_PRETTY_PRINT);
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

                $user = $this->gateway->getByTitre($data["titre"]);

                if ($user) {
                    http_response_code(404);
                    echo json_encode(["message" => "titre already used "], JSON_PRETTY_PRINT);
                    break;
                }

                $user = $this->gateway->getByDescription($data["description"]);

                if ($user) {
                    http_response_code(404);
                    echo json_encode(["message" => "description already used "], JSON_PRETTY_PRINT);
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

}
