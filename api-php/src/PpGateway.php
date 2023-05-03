<?php

class PpGateway
{
    private PDO $conn;

    public function __construct(Database $database)
    {
        $this->conn = $database->getConnection();
    }

    public function getAll(): array
    {
        $sql = "SELECT * 
            FROM imagepp";

        $stmt = $this->conn->query($sql);

        $data = [];

        while ($row = $stmt->fetch(PDO::FETCH_ASSOC)) {
            $data[] = $row;
        }

        return $data;

    }

    public function get(string $id): array|false
    {
        $sql = "SELECT *
                FROM imagepp
                WHERE id_pp = :id_pp";

        $stmt = $this->conn->prepare($sql);

        $stmt->bindValue(":id_pp", htmlspecialchars($id), PDO::PARAM_INT);

        $stmt->execute();

        $data = $stmt->fetch(PDO::FETCH_ASSOC);

        return $data;
    }

}