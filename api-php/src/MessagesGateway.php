<?php

class MessagesGateway
{
    private PDO $conn;

    public function __construct(Database $database)
    {
        $this->conn = $database->getConnection();
    }

    //récupère tous les messages
    public function getAll(): array
    {
        $sql = "SELECT * 
            FROM messages";

        $stmt = $this->conn->query($sql);

        $data = [];

        while ($row = $stmt->fetch(PDO::FETCH_ASSOC)) {
            $data[] = $row;
        }

        return $data;

    }

    //créer un message
    public function create(array $data): string
    {
        $sql = "INSERT INTO messages (message, id_user, publi_time, id_topics) 
                VALUES (:message, :id_user, NOW(), :id_topics)";

        $stmt = $this->conn->prepare($sql);
        if (isset($data["message"])) {
            $data["message"] = htmlspecialchars($data["message"]);
        }

        $stmt->bindValue(":message", $data["message"], PDO::PARAM_STR);
        $stmt->bindValue(":id_user", htmlspecialchars($data["id_user"]), PDO::PARAM_INT);
        $stmt->bindValue(":id_topics", htmlspecialchars($data["id_topics"]), PDO::PARAM_INT);

        $stmt->execute();

        return $this->conn->lastInsertId();

    }

    //récupérer un message par son id
    public function get(string $id): array|false
    {
        $sql = "SELECT *
                FROM messages
                WHERE id_message = :id_message";

        $stmt = $this->conn->prepare($sql);

        $stmt->bindValue(":id_message", htmlspecialchars($id), PDO::PARAM_INT);

        $stmt->execute();

        $data = $stmt->fetch(PDO::FETCH_ASSOC);

        return $data;
    }

    //récupérer des messages par l'id de leur topics
    public function getByTopics(string $topics): array|false
    {
        $sql = "SELECT *
                FROM messages
                WHERE id_topics = :id_topics";

        $stmt = $this->conn->prepare($sql);

        $stmt->bindValue(":id_topics", htmlspecialchars($topics), PDO::PARAM_INT);

        $stmt->execute();

        $data = [];

        while ($row = $stmt->fetch(PDO::FETCH_ASSOC)) {
            $data[] = $row;
        }

        return $data;
    }

    //modifier un message par son id
    public function update(array $current, array $new): int
    {
        $sql = "UPDATE messages
                SET message = :message, id_user = :id_user, publi_time= :publi_time, id_topics = :id_topics
                WHERE id_message = :id_message";

        $stmt = $this->conn->prepare($sql);

        if (isset($new["message"])) {
            $new["message"] = htmlspecialchars($new["message"]);
        }
        if (isset($new["id_user"])) {
            $new["id_user"] = htmlspecialchars($new["id_user"]);
        }
        if (isset($new["publi_time"])) {
            $new["publi_time"] = htmlspecialchars($new["publi_time"]);
        }
        if (isset($new["id_topics"])) {
            $new["id_topics"] = htmlspecialchars($new["id_topics"]);
        }

        $stmt->bindValue(
            ":message", $new["message"] ?? $current["message"],
            PDO::PARAM_STR
        );
        $stmt->bindValue(
            ":id_user", $new["id_user"] ?? $current["id_user"],
            PDO::PARAM_INT
        );
        $stmt->bindValue(
            ":publi_time", $new["publi_time"] ?? $current["publi_time"],
            PDO::PARAM_STR
        );
        $stmt->bindValue(
            ":id_topics", $new["id_topics"] ?? $current["id_topics"],
            PDO::PARAM_INT
        );

        $stmt->bindValue(":id_message", $current["id_message"], PDO::PARAM_INT);

        $stmt->execute();

        return $stmt->rowCount();
    }

    //suprimer un message par son id
    public function delete(string $id): int
    {
        $sql = "DELETE FROM messages
                WHERE id_message = :id_message";

        $stmt = $this->conn->prepare($sql);

        $stmt->bindValue(":id_message", htmlspecialchars($id), PDO::PARAM_INT);

        $stmt->execute();

        return $stmt->rowCount();
    }

}