<?php

class TopicsGateway
{
    private PDO $conn;

    public function __construct(Database $database)
    {
        $this->conn = $database->getConnection();
    }

    public function getAll(): array
    {
        $sql = "SELECT * 
            FROM topics";

        $stmt = $this->conn->query($sql);

        $data = [];

        while ($row = $stmt->fetch(PDO::FETCH_ASSOC)) {
            $data[] = $row;
        }

        return $data;

    }

    public function create(array $data): string
    {
        $sql = "INSERT INTO topics (titre, description, crea_date, id_tags, id_user) 
                VALUES (:titre, :description, NOW(), :id_tags, :id_user)";

        $stmt = $this->conn->prepare($sql);

        $stmt->bindValue(":titre", htmlspecialchars($data["titre"]), PDO::PARAM_STR);
        $stmt->bindValue(":description", htmlspecialchars($data["description"]), PDO::PARAM_STR);
        $stmt->bindValue(":id_tags", htmlspecialchars($data["id_tags"]), PDO::PARAM_INT);
        $stmt->bindValue(":id_user", htmlspecialchars($data["id_user"]), PDO::PARAM_INT);

        $stmt->execute();

        return $this->conn->lastInsertId();

    }

    public function get(string $id): array|false
    {
        $sql = "SELECT *
                FROM topics
                WHERE id_topics = :id_topics";

        $stmt = $this->conn->prepare($sql);

        $stmt->bindValue(":id_topics", htmlspecialchars($id), PDO::PARAM_INT);

        $stmt->execute();

        $data = $stmt->fetch(PDO::FETCH_ASSOC);

        return $data;
    }

    public function getByTags(string $tags): array
    {
        $sql = "SELECT *
                FROM topics
                WHERE id_tags = :id_tags";

        $stmt = $this->conn->prepare($sql);

        $stmt->bindValue(":id_tags", htmlspecialchars($tags), PDO::PARAM_INT);

        $stmt->execute();

        $data = [];

        while ($row = $stmt->fetch(PDO::FETCH_ASSOC)) {
            $data[] = $row;
        }

        return $data;
    }

    public function getByTitre(string $titre): array|false
    {
        $sql = "SELECT *
                FROM topics
                WHERE titre = :titre";

        $stmt = $this->conn->prepare($sql);

        $stmt->bindValue(":titre", htmlspecialchars($titre), PDO::PARAM_STR);

        $stmt->execute();

        $data = $stmt->fetch(PDO::FETCH_ASSOC);

        return $data;
    }

    public function getByDescription(string $description): array|false
    {
        $sql = "SELECT *
                FROM topics
                WHERE description = :description";

        $stmt = $this->conn->prepare($sql);

        $stmt->bindValue(":description", htmlspecialchars($description), PDO::PARAM_STR);

        $stmt->execute();

        $data = $stmt->fetch(PDO::FETCH_ASSOC);

        return $data;
    }

}