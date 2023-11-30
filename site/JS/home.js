import { User } from "/site/JS/user_class.js";
import { Imagepp } from "/site/JS/pp_class.js";
import { Tags } from "/site/JS/tags_class.js";
import { Topics } from "/site/JS/topics_class.js";
import { Messages } from "/site/JS/messages_class.js";

let api_port

// charge le site avec l
$.getJSON("/site/conf.json", function (data) {
    api_port = data.api_port
    fetch_tags();
    fetch_all();
    post_tag();
    if(localUser){
        get_pp()
    }
})

const CreateMessage = document.getElementById("CreateMessage");
const favDialog = document.getElementById("favDialog");
const swmode = document.getElementById("swmode");
const localUser = localStorage.getItem("loged_user")?.toString();
let storageUser = new User("", "", "", "", "", "");

if (localUser) {
    console.log("auto connect");
    storageUser = JSON.parse(localUser);

    switch (storageUser.theme) {
        case "dark":
            document.querySelector("body").setAttribute("data-theme", "dark");
            console.log("default dark");
            swmode.checked = true;
            break;
        case "light":
            document.querySelector("body").setAttribute("data-theme", "light");
            console.log("default light");
            swmode.checked = false;
            break;
    }
} else {
    console.log("to connect");
    CreateMessage.style.display = "none"
}

const btn_deconenxion = document.getElementById("btn_deconenxion");

btn_deconenxion.addEventListener("click", () => {
    localStorage.removeItem("loged_user");
    window.location.href = "/Forum/home";
});

let publisher;
let pp_publi;
let list_tags = [];
let list_topics = [];
let list_messages = [];
const display_tags = document.getElementById("display_tags");
const display_topics = document.getElementById("display_topics");
const style_mod = document.getElementById("style_mod");

async function fetch_all() {
    favDialog.style.display = "none";
    display_topics.style.position = "inherit";
    const topicsload = await fetch(window.location.origin + "/apiForum/topics", {
        method: "GET",
        headers: {
            Accept: "application/json",
            "Content-type": "application/json; charset=UTF-8",
        },
    }).then((res) => {
        if (res.ok) {
            res.json().then((data) => {
                display_topics.innerHTML = "";
                list_topics = []
                data.forEach((elt) => {
                    let time = new Date(elt.crea_date)
                    elt.crea_date = time.toLocaleDateString('fr') + " à " + time.toLocaleTimeString("fr")
                    let actual_topic = new Topics(
                        elt.id_topics,
                        elt.titre,
                        elt.description,
                        elt.crea_date,
                        elt.id_tags,
                        elt.id_user
                    );
                    list_topics.push(actual_topic);

                });

                list_topics.forEach(elt => {

                    const topicsuserload = fetch(
                        window.location.origin + `/apiForum/users/${elt.id_user}`,
                        {
                            method: "GET",
                            headers: {
                                Accept: "application/json",
                                "Content-type": "application/json; charset=UTF-8",
                            },
                        }
                    ).then((res) => {
                        if (res.ok) {
                            res.json().then((data) => {
                                publisher = new User(
                                    data.id_user,
                                    data.pseudo,
                                    data.email,
                                    data.passwd,
                                    data.id_imagepp,
                                    data.theme
                                );

                                display_topics.innerHTML += `
                                        <div class="card" id="topics${elt.id_topics}" onclick=openmessage(${elt.id_topics})>
                                            <div class="top_card">
                                                <h4 class="user_card${publisher.id_imagepp}">${publisher.pseudo}</h4>
                                                <p> &ensp; publié le ${elt.crea_date}</p>
                                            </div>
                                            <div class="middle_card">
                                                <h3 class="title_topic${elt.id_tags}">${elt.titre}</h3>
                                            </div>
                                            <div class="bottom_card">
                                                <p>${elt.description}</p>
                                            </div>
                                        </div>`;

                                style_mod.innerHTML += `
            .title_topic${elt.id_tags}::before {
            content: url(/site/Assets/Images/icon_tag/tags${elt.id_tags}.svg);
            }
            .title_topic${elt.id_tags} {
            text-align: center;
            }`;

                                const ppload = fetch(
                                    window.location.origin + `/apiForum/pp/${publisher.id_imagepp}`,
                                    {
                                        method: "GET",
                                        headers: {
                                            Accept: "application/json",
                                            "Content-type": "application/json; charset=UTF-8",
                                        },
                                    }
                                ).then((res) => {
                                    if (res.ok) {
                                        res.json().then((data) => {
                                            pp_publi = new Imagepp(data.id_pp, data.image_loc);

                                            style_mod.innerHTML += `
            .user_card${data.id_pp}::before {
            content: url(/site/Assets/Images/profil/${pp_publi.image_loc});
            }`;
                                        });
                                    } else {
                                        console.log("res.ok false");
                                    }
                                });
                            });
                        } else {
                            console.log("res.ok false");
                        }
                    });
                })


            });
        } else {
            console.log("res.ok false");
        }
    });
}

export function openmessage(id_topic) {
    fetch_by_topics(id_topic);
    fetch_by_topics_messages(id_topic);
    selectedTopic = id_topic
    favDialog.style.display = "block";
    display_topics.style.position = "fixed";
    popup_create.style.display = "none";
}

let selectedTopic = 0;

async function fetch_tags() {
    const tagsload = await fetch("http://localhost:" + api_port + "/apiForum/tags", {
        method: "GET",
        headers: {
            Accept: "application/json",
            "Content-type": "application/json; charset=UTF-8",
        },
    }).then((res) => {
        if (res.ok) {
            res.json().then((data) => {
                data.forEach((elt) => {
                    list_tags.push(new Tags(elt.id_tags, elt.tags));
                });

                display_tags.innerHTML += `<h4 class="tag tagsall" id="tagsall"><span class="hover-underline-animation">All Tags</span></h4>`;
                style_mod.innerHTML += `
.tagsall::before{
    content: url(/site/Assets/Images/icon_input/arobase.svg);
    position: absolute;
    left: -20px;
}`;

                list_tags.forEach((elt) => {
                    display_tags.innerHTML += `<h4 class="tag tags${elt.id_tags}" id="tags${elt.id_tags}" style= ""><span class="hover-underline-animation">${elt.tags}</span></h4>`;
                    style_mod.innerHTML += `
.tags${elt.id_tags}::before{
    content: url(/site/Assets/Images/icon_tag/tags${elt.id_tags}.svg);
    position: absolute;
    left: -20px;
}`;
                });

                document.getElementById(`tagsall`).onclick = function () {
                    fetch_all();
                };

                list_tags.forEach((elt) => {
                    document.getElementById(`tags${elt.id_tags}`).onclick = function () {
                        fetch_by_tags(`${elt.id_tags}`);
                    };
                });
            });
        } else {
            console.log("res.ok false");
        }
    });
}

async function fetch_by_tags(tag) {
    favDialog.style.display = "none";
    display_topics.style.position = "inherit";
    const topicsload = await fetch(
        window.location.origin + `/apiForum/topicstags/${tag}`,
        {
            method: "GET",
            headers: {
                Accept: "application/json",
                "Content-type": "application/json; charset=UTF-8",
            },
        }
    ).then((res) => {
        if (res.ok) {
            res.json().then((datas) => {
                display_topics.innerHTML = "";
                list_topics = []
                datas.forEach((elt) => {
                    let time = new Date(elt.crea_date)
                    elt.crea_date = time.toLocaleDateString('fr') + " à " + time.toLocaleTimeString("fr")
                    let actual_topic = new Topics(
                        elt.id_topics,
                        elt.titre,
                        elt.description,
                        elt.crea_date,
                        elt.id_tags,
                        elt.id_user
                    );
                    list_topics.push(actual_topic);

                });

                list_topics.forEach(elt => {

                    const topicsload = fetch(
                        window.location.origin + `/apiForum/users/${elt.id_user}`,
                        {
                            method: "GET",
                            headers: {
                                Accept: "application/json",
                                "Content-type": "application/json; charset=UTF-8",
                            },
                        }
                    ).then((res) => {
                        if (res.ok) {
                            res.json().then((data) => {
                                publisher = new User(
                                    data.id_user,
                                    data.pseudo,
                                    data.email,
                                    data.passwd,
                                    data.id_imagepp,
                                    data.theme
                                );

                                const topicsload = fetch(window.location.origin + `/apiForum/users/${elt.id_user}`, {
                                    method: 'GET',
                                    headers: {
                                        "Accept": "application/json",
                                        "Content-type": "application/json; charset=UTF-8"
                                    }
                                })
                                    .then((res) => {
                                        if (res.ok) {
                                            res.json().then(data => {
                                                publisher = new User(data.id_user, data.pseudo, data.email, data.passwd, data.id_imagepp, data.theme)

                                                display_topics.innerHTML += `
                                        <div class="card"  id="topics${elt.id_topics}" onclick=openmessage(${elt.id_topics})>
                                            <div class="top_card">
                                                <h4 class="user_card${publisher.id_imagepp}">${publisher.pseudo}</h4>
                                                <p> &ensp; publié le ${elt.crea_date}</p>
                                            </div>
                                            <div class="middle_card">
                                                <h3 class="title_topic${elt.id_tags}">${elt.titre}</h3>
                                            </div>
                                            <div class="bottom_card">
                                                <p>${elt.description}</p>
                                            </div>
                                        </div>`;

                                                style_mod.innerHTML += `
.title_topic${elt.id_tags}::before {
content: url(/site/Assets/Images/icon_tag/tags${elt.id_tags}.svg);
}
.title_topic${elt.id_tags} {
text-align: center;
}`;

                                                const ppload = fetch(
                                                    window.location.origin + `/apiForum/pp/${publisher.id_imagepp}`,
                                                    {
                                                        method: "GET",
                                                        headers: {
                                                            Accept: "application/json",
                                                            "Content-type": "application/json; charset=UTF-8",
                                                        },
                                                    }
                                                ).then((res) => {
                                                    if (res.ok) {
                                                        res.json().then((data) => {
                                                            pp_publi = new Imagepp(data.id_pp, data.image_loc);

                                                            style_mod.innerHTML += `
.user_card${data.id_pp}::before {
content: url(../../site/Assets/Images/profil/${pp_publi.image_loc});
}`;
                                                        });
                                                    } else {
                                                        console.log("res.ok false");
                                                    }
                                                });
                                            });
                                        } else {
                                            console.log("res.ok false");
                                        }
                                    });
                            });
                        } else {
                            console.log("res.ok false");
                        }
                    });
                })
            });
        } else {
            console.log("res.ok false");
            display_topics.innerHTML = `<h3>Aucun topic trouvé</h3>`
        }
    });
}

const left_container = document.getElementById("left_container");

async function fetch_by_topics(id_topics) {
    const topicsload = await fetch(
        window.location.origin + `/apiForum/topics/${id_topics}`,
        {
            method: "GET",
            headers: {
                Accept: "application/json",
                "Content-type": "application/json; charset=UTF-8",
            },
        }
    ).then((res) => {
        if (res.ok) {
            res.json().then((data) => {
                let time = new Date(data.crea_date)
                data.crea_date = time.toLocaleDateString('fr') + " à " + time.toLocaleTimeString("fr")
                let actual_topic = new Topics(
                    data.id_topics,
                    data.titre,
                    data.description,
                    data.crea_date,
                    data.id_tags,
                    data.id_user
                );

                const userload = fetch(
                    window.location.origin + `/apiForum/users/${data.id_user}`,
                    {
                        method: "GET",
                        headers: {
                            Accept: "application/json",
                            "Content-type": "application/json; charset=UTF-8",
                        },
                    }
                ).then((res) => {
                    if (res.ok) {
                        res.json().then((data) => {
                            let publisher = new User(
                                data.id_user,
                                data.pseudo,
                                data.email,
                                data.passwd,
                                data.id_imagepp,
                                data.theme
                            );
                            left_container.innerHTML = "";
                            left_container.innerHTML += `
                    <div class="post">
                        <div class="top_post">
                                <div class="info-top_post">
                                    <h4 class="user_card_post_comment user_card${publisher.id_imagepp}">${publisher.pseudo}</h4>
                                    <p> &ensp; publié le ${actual_topic.crea_date}</p>
                                </div>
                            </div>
                            <div class="title_post">
                                <h3 class="title_topic${actual_topic.id_tags}">${actual_topic.titre}</h3>
                            </div>
                            <div class="main_post">
                                <p>${actual_topic.description}</p>
                            </div>
                        </div>
                    </div>`;

                            style_mod.innerHTML += `
                    .title_topic${actual_topic.id_tags}::before {
                        content: url(/site/Assets/Images/icon_tag/tags${actual_topic.id_tags}.svg);
                    }
                    .title_topic${actual_topic.id_tags} {
                      text-align: center;
                    }`;

                            const ppload = fetch(
                                window.location.origin + `/apiForum/pp/${publisher.id_imagepp}`,
                                {
                                    method: "GET",
                                    headers: {
                                        Accept: "application/json",
                                        "Content-type": "application/json; charset=UTF-8",
                                    },
                                }
                            ).then((res) => {
                                if (res.ok) {
                                    res.json().then((data) => {
                                        let pp_publi = new Imagepp(data.id_pp, data.image_loc);

                                        style_mod.innerHTML += `
                          .user_card${data.id_pp}::before {
                            content: url(/site/Assets/Images/profil/${pp_publi.image_loc});
                          }`;
                                    });
                                } else {
                                    console.log("res.ok false");
                                }
                            });
                        });
                    } else {
                        console.log("res.ok false");
                    }
                });
            });
        } else {
            console.log("res.ok false");
        }
    });
}

const comment_container = document.getElementById("comment_container");

async function fetch_by_topics_messages(id_topics) {
    const messagesload = await fetch(
        window.location.origin + `/apiForum/messagestopics/${id_topics}`,
        {
            method: "GET",
            headers: {
                Accept: "application/json",
                "Content-type": "application/json; charset=UTF-8",
            },
        }
    ).then((res) => {
        if (res.ok) {
            res.json().then((data) => {
                comment_container.innerHTML = "";
                list_messages = []
                data.forEach((elt) => {
                    let time = new Date(elt.publi_time)
                    elt.publi_time = time.toLocaleDateString('fr') + " à " + time.toLocaleTimeString("fr")

                    let actual_message = new Messages(
                        elt.id_message,
                        elt.message,
                        elt.id_user,
                        elt.publi_time,
                        elt.id_topics
                    );
                    list_messages.push(actual_message);

                });

                list_messages.forEach(elt => {

                    const userload = fetch(
                        window.location.origin + `/apiForum/users/${elt.id_user}`,
                        {
                            method: "GET",
                            headers: {
                                Accept: "application/json",
                                "Content-type": "application/json; charset=UTF-8",
                            },
                        }
                    ).then((res) => {
                        if (res.ok) {
                            res.json().then((data) => {
                                let publisher = new User(
                                    data.id_user,
                                    data.pseudo,
                                    data.email,
                                    data.passwd,
                                    data.id_imagepp,
                                    data.theme
                                );

                                comment_container.innerHTML += `
                                <div class="comment_user">
                                    <div class="comment_user_name">
                                        <h4 class="user_card${publisher.id_imagepp}">${publisher.pseudo}</h4>
                                        <p> &ensp; publié le ${elt.publi_time}</p>
                                    </div>
                                    <div class="comment_user_text">
                                        <p>${elt.message}</p>
                                    </div>
                                </div>`;

                                const ppload = fetch(
                                    window.location.origin + `/apiForum/pp/${publisher.id_imagepp}`,
                                    {
                                        method: "GET",
                                        headers: {
                                            Accept: "application/json",
                                            "Content-type": "application/json; charset=UTF-8",
                                        },
                                    }
                                ).then((res) => {
                                    if (res.ok) {
                                        res.json().then((data) => {
                                            let pp_publi = new Imagepp(data.id_pp, data.image_loc);

                                            style_mod.innerHTML += `
                  .user_card${data.id_pp}::before {
                    content: url(/site/Assets/Images/profil/${pp_publi.image_loc});
                  }`;
                                        });
                                    } else {
                                        console.log("res.ok false");
                                    }
                                });
                            });
                        } else {
                            console.log("res.ok false");
                        }
                    });
                })

            });
        } else {
            console.log("res.ok false");
            comment_container.innerHTML = `<h3>Aucun message trouvé</h3>`
        }
    });
}

CreateMessage.addEventListener("submit", async (event) => {
    event.preventDefault();

    const message = document.getElementById("comment").value;
    document.getElementById("comment").value = ""
    const userId = storageUser.id_user;
    const topics = selectedTopic;

    if (message != "") {

        const response = await fetch("http://localhost:" + api_port + "/apiForum/messages", {
            method: "POST",
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json; charset=UTF-8",
            },
            body: JSON.stringify({
                message: message,
                id_topics: parseInt(topics),
                id_user: userId,
            }),
        })
        setTimeout(function () {
            fetch_by_topics_messages(selectedTopic)
        }, 400)
    }
});


const tagDropdown = document.getElementById("tag-dropdown");

async function post_tag() {
    const tagsload = await fetch("http://localhost:" + api_port + "/apiForum/tags", {
        method: "GET",
        headers: {
            Accept: "application/json",
            "Content-type": "application/json; charset=UTF-8",
        },
    })
        .then((response) => response.json())
        .then((data) => {
            data.forEach((tag) => {
                const option = document.createElement("option");
                option.value = tag.id_tags;
                option.text = tag.tags;
                tagDropdown.appendChild(option);
            });
        })
        .catch((error) => console.error(error));
}

const form = document.getElementById("Createpost");
const statusMessage = document.getElementById("status-message");

form.addEventListener("submit", async (event) => {
    event.preventDefault();

    const titre = document.getElementById("post-title").value;
    const userId = storageUser.id_user;
    const description = document.getElementById("post-description").value;
    const selectedTag = tagDropdown.value;

    if (selectedTag == "") {
        statusMessage.textContent = "Choisissez un tag";
        return;
    }

    const response = await fetch("http://localhost:" + api_port + "/apiForum/topics", {
        method: "POST",
        headers: {
            "Accept": "application/json",
            "Content-Type": "application/json; charset=UTF-8"
        },
        body: JSON.stringify({ titre: titre, description: description, id_tags: parseInt(selectedTag), id_user: parseInt(userId) })
    });

    if (response.ok) {
        statusMessage.textContent = `Message ajouté avec succès`;
        setTimeout(function () {
            popup_create.style.display = "none";
            fetch_by_tags(selectedTag)
        }, 2000)
    } else {
        statusMessage.textContent = "Erreur lors de l'ajout du message";
    }
});

const myBtn = document.getElementById("myBtn");
const popup_create = document.querySelector(".popup_create");
const post_close = document.querySelector(".post_close");
const profil_pseudo = document.getElementById("profil_pseudo");

if (localUser) {
    // When the user clicks on the button, open the modal
    myBtn.onclick = function () {
        popup_create.style.display = "block";
        display_topics.style.position = "fixed";
        favDialog.style.display = "none";
    };

    // When the user clicks on <span> (x), close the modal
    post_close.onclick = function () {
        popup_create.style.display = "none";
        display_topics.style.position = "inherit";
    };

    // When the user clicks anywhere outside of the modal, close it
    window.onclick = function (event) {
        if (event.target == popup_create) {
            popup_create.style.display = "none";
        }
    };
} else {
    myBtn.onclick = function () {
        document.location.href = "/Forum/log_in";
    };
}

const div_text = document.querySelector(".div_text");
const div_text_connect = document.querySelector(".div_text_connect");

let profil_picture = document.getElementById("profil_picture")

async function get_pp() {
    const tagsload = await fetch("http://localhost:" + api_port + "/apiForum/pp/" + storageUser.id_imagepp, {
        method: "GET",
        headers: {
            Accept: "application/json",
            "Content-type": "application/json; charset=UTF-8",
        },
    })
        .then((response) => response.json())
        .then((data) => {
            profil_picture.src = "/site/Assets/Images/profil/"+data.image_loc
        })
        .catch((error) => console.error(error));
}

function log_In() {
    if (localUser) {
        open_create.style.display = "none";
        div_text.style.display = "none";
        div_text_connect.style.display = "block";
        profil_pseudo.innerHTML = storageUser.pseudo;
        display_topics.style.position = "inherit";
    } else {
        display_topics.style.position = "fixed";
        open_create.style.display = "block";
        div_text.style.display = "block";
        div_text_connect.style.display = "none";
    }
}

log_In();

open_create.addEventListener("click", function () {
    document.location.href = "/Forum/log_in";
});

//pop-up commentaire
const close_pop = document.getElementById("close_pop");

close_pop.addEventListener("click", function () {
    favDialog.style.display = "none";
    display_topics.style.position = "inherit";
});
