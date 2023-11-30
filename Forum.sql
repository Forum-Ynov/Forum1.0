
-- Listage de la structure de table kcdxbd_sqwadowe_db. imagepp
CREATE TABLE IF NOT EXISTS `imagepp` (
  `id_pp` int(10) NOT NULL AUTO_INCREMENT,
  `image_loc` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id_pp`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

-- Listage des données de la table kcdxbd_sqwadowe_db.imagepp : ~6 rows (environ)
INSERT INTO `imagepp` (`id_pp`, `image_loc`) VALUES
	(1, 'anonymous.svg'),
	(2, 'chat.svg'),
	(3, 'hatsune-miku.svg'),
	(4, 'homer.svg'),
	(5, 'koala.svg'),
	(6, 'mouton.svg');

-- Listage de la structure de table kcdxbd_sqwadowe_db. messages
CREATE TABLE IF NOT EXISTS `messages` (
  `id_message` int(10) NOT NULL AUTO_INCREMENT,
  `message` varchar(255) NOT NULL DEFAULT '',
  `id_user` int(10) NOT NULL DEFAULT 1,
  `publi_time` datetime NOT NULL,
  `id_topics` int(10) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id_message`),
  KEY `FK_messages_user` (`id_user`),
  KEY `FK_messages_topics` (`id_topics`),
  CONSTRAINT `FK_messages_topics` FOREIGN KEY (`id_topics`) REFERENCES `topics` (`id_topics`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `FK_messages_user` FOREIGN KEY (`id_user`) REFERENCES `user` (`id_user`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

-- Listage des données de la table kcdxbd_sqwadowe_db.messages : ~8 rows (environ)
INSERT INTO `messages` (`id_message`, `message`, `id_user`, `publi_time`, `id_topics`) VALUES
	(1, 'api + site = cool', 16, '2023-04-20 23:31:08', 1),
	(2, 'Navigateur mobile ', 17, '2023-04-21 11:23:13', 2),
	(3, 'pok', 16, '2023-05-03 22:10:18', 2),
	(21, 'Bonjour toute le monde, helloo guys, je suis moi', 17, '2023-05-03 22:51:37', 3),
	(44, 'test', 16, '2023-05-04 01:06:24', 2),
	(45, 'cool', 16, '2023-05-04 10:09:55', 2),
	(46, 'truc', 17, '2023-05-04 10:10:35', 1),
	(47, 'nul', 16, '2023-05-04 14:08:33', 6),
	(48, '&lt;script&gt; alert(&#34;hacked&#34;) &lt;/script&gt;', 16, '2023-05-04 15:38:18', 6),
	(49, 'dhddydrdwdg', 16, '2023-05-04 16:30:10', 8);

-- Listage de la structure de table kcdxbd_sqwadowe_db. tags
CREATE TABLE IF NOT EXISTS `tags` (
  `id_tags` int(10) NOT NULL AUTO_INCREMENT,
  `tags` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id_tags`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

-- Listage des données de la table kcdxbd_sqwadowe_db.tags : ~5 rows (environ)
INSERT INTO `tags` (`id_tags`, `tags`) VALUES
	(1, 'Life Style'),
	(2, 'Jeu Vidéo'),
	(3, 'Cuisine'),
	(4, 'Dev/Code'),
	(5, 'Football');

-- Listage de la structure de table kcdxbd_sqwadowe_db. topics
CREATE TABLE IF NOT EXISTS `topics` (
  `id_topics` int(10) NOT NULL AUTO_INCREMENT,
  `titre` varchar(255) NOT NULL DEFAULT '',
  `description` varchar(255) NOT NULL DEFAULT '',
  `crea_date` datetime NOT NULL,
  `id_tags` int(10) NOT NULL DEFAULT 1,
  `id_user` int(10) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id_topics`),
  KEY `FK_topics_tags` (`id_tags`),
  KEY `FK_topics_user` (`id_user`),
  CONSTRAINT `FK_topics_tags` FOREIGN KEY (`id_tags`) REFERENCES `tags` (`id_tags`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `FK_topics_user` FOREIGN KEY (`id_user`) REFERENCES `user` (`id_user`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

-- Listage des données de la table kcdxbd_sqwadowe_db.topics : ~7 rows (environ)
INSERT INTO `topics` (`id_topics`, `titre`, `description`, `crea_date`, `id_tags`, `id_user`) VALUES
	(1, 'enfin en ligne', 'on est chez webstrator', '2023-04-20 23:30:44', 4, 16),
	(2, 'Nouveau sujet sur mobile', 'À été créé sur mobile ', '2023-04-21 11:22:11', 1, 16),
	(3, 'test avec enzo', 'il y a ausii dylan', '2023-04-24 08:26:01', 5, 17),
	(4, 'kkzerlzjitefo', ',lks,fklznlfef', '2023-04-24 10:13:51', 2, 16),
	(5, 'pour le scroll', 'pour le scroll', '2023-04-24 10:14:19', 5, 17),
	(6, 'test avant oral', 'c&#39;est la merde lol (c\'est faux)', '2023-05-04 14:08:17', 4, 16),
	(7, 'ok', 'ok', '2023-05-04 15:45:28', 2, 16),
	(8, 'xwvsg', 'rdreh', '2023-05-04 16:29:58', 3, 16);

-- Listage de la structure de table kcdxbd_sqwadowe_db. user
CREATE TABLE IF NOT EXISTS `user` (
  `id_user` int(10) NOT NULL AUTO_INCREMENT,
  `pseudo` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `passwd` varchar(255) NOT NULL DEFAULT '',
  `id_imagepp` int(10) NOT NULL DEFAULT 1,
  `theme` varchar(255) NOT NULL DEFAULT 'light',
  PRIMARY KEY (`id_user`),
  KEY `id_imagepp` (`id_imagepp`),
  CONSTRAINT `imagepp` FOREIGN KEY (`id_imagepp`) REFERENCES `imagepp` (`id_pp`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

-- Listage des données de la table kcdxbd_sqwadowe_db.user : ~2 rows (environ)
INSERT INTO `user` (`id_user`, `pseudo`, `email`, `passwd`, `id_imagepp`, `theme`) VALUES
	(16, 'mateo', 'mateoluque@aol.com', '$2a$14$aYiMig8XzdQ/vSndSlvxLO8IZWbQP0LEg5Rp9AkLs7Lq2dna0aUAO', 5, 'dark'),
	(17, 'lorenzo', 'lorenzo.gotti@ynov.com', '$2a$14$yUVCzWRozD4BrBXl7Hp97.hFqVeBDJD.hOgzzgAFnUnvfsgFJSvta', 2, 'dark');
