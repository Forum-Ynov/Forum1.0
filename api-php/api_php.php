<?php

declare(strict_types=1);

require_once realpath(__DIR__ . '/vendor/autoload.php');
$dotenv = Dotenv\Dotenv::createImmutable(__DIR__);
$dotenv->load();

spl_autoload_register(function ($class) {
    require __DIR__ . "/src/$class.php";
});

set_error_handler("ErrorHandler::handleError");
set_exception_handler("ErrorHandler::handleException");

header("Content-Type: application/json; charset=UTF-8");

// required headers
header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Headers: Origin, X-Requested-With, Content-Type, Accept, Authorization');
header("Access-Control-Allow-Credentials: true");
header("Access-Control-Allow-Methods: *");
header('Content-Type: application/json');
$method = $_SERVER['REQUEST_METHOD'];
if ($method == "OPTIONS") {
    header('Access-Control-Allow-Origin: *');
    header("Access-Control-Allow-Headers: X-API-KEY, Origin, X-Requested-With, Content-Type, Accept, Access-Control-Request-Method,Access-Control-Request-Headers, Authorization");
    header("HTTP/1.1 200 OK");
    die();
}

$parts = explode("/", $_SERVER["REQUEST_URI"]);

$database = new Database($_ENV["DB_HOST"], $_ENV["DB_PORT"], $_ENV["DB_DATABASE"], $_ENV["DB_USER"], $_ENV["DB_PASSWORD"]);

if ($parts[1] == "apiForum") {
    if (isset($parts[2])) {


        if ($parts[2] == "users") {
            $id = $parts[3] ?? null;
            $gateway = new UserGateway($database);
            $controller = new UserController($gateway);

            $controller->processRequest($_SERVER["REQUEST_METHOD"], $id);

        } elseif ($parts[2] == "userpseudo") {
            $pseudo = $parts[3] ?? null;
            $gateway = new UserGateway($database);
            $controller = new UserController($gateway);

            $controller->processRequestByPseudo($_SERVER["REQUEST_METHOD"], $pseudo);

        } elseif ($parts[2] == "useremail") {
            $email = $parts[3] ?? null;
            $gateway = new UserGateway($database);
            $controller = new UserController($gateway);

            $controller->processRequestByEmail($_SERVER["REQUEST_METHOD"], $email);

        } elseif ($parts[2] == "login") {
            $gateway = new UserGateway($database);
            $controller = new UserController($gateway);

            $controller->processRequestLogin($_SERVER["REQUEST_METHOD"]);

        } elseif ($parts[2] == "pp") {
            $id = $parts[3] ?? null;
            $gateway = new PpGateway($database);
            $controller = new PpController($gateway);

            $controller->processRequest($_SERVER["REQUEST_METHOD"], $id);

        } elseif ($parts[2] == "tags") {
            $id = $parts[3] ?? null;
            $gateway = new TagsGateway($database);
            $controller = new TagsController($gateway);

            $controller->processRequest($_SERVER["REQUEST_METHOD"], $id);

        } elseif ($parts[2] == "topics") {
            $id = $parts[3] ?? null;
            $gateway = new TopicsGateway($database);
            $controller = new TopicsController($gateway);

            $controller->processRequest($_SERVER["REQUEST_METHOD"], $id);

        } elseif ($parts[2] == "topicstags") {
            $tags = $parts[3] ?? null;
            $gateway = new TopicsGateway($database);
            $controller = new TopicsController($gateway);

            $controller->processRequestByTags($_SERVER["REQUEST_METHOD"], $tags);

        } elseif ($parts[2] == "messages") {
            $id = $parts[3] ?? null;
            $gateway = new MessagesGateway($database);
            $controller = new MessagesController($gateway);

            $controller->processRequest($_SERVER["REQUEST_METHOD"], $id);

        } elseif ($parts[2] == "messagestopics") {
            $topics = $parts[3] ?? null;
            $gateway = new MessagesGateway($database);
            $controller = new MessagesController($gateway);

            $controller->processRequestByTopics($_SERVER["REQUEST_METHOD"], $topics);

        } else {
            http_response_code(404);
            exit;
        }
    }
}

?>