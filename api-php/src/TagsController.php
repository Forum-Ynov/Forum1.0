<?php

class TagsController
{
    public function __construct(private TagsGateway $gateway)
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

    private function processRessourceRequest(string $method, string $id): void
    {
        $pp = $this->gateway->get($id);

        if (!$pp) {
            http_response_code(404);
            echo json_encode(["message" => "Tags not found"], JSON_PRETTY_PRINT);
            return;
        }

        switch ($method) {
            case "GET":
                echo json_encode($pp, JSON_PRETTY_PRINT);
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
            
            default:
                http_response_code(405);
                header("Allow: GET");
        }
    }

}

