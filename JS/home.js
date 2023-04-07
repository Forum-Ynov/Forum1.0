
import { User } from "./user_class.js";
import { Imagepp } from "./pp_class.js";
import { Tags } from "./tags_class.js";
import { Topics } from "./topics_class.js";
import { Messages } from "./messages_class.js";

class Cards {
    constructor(Topics, Tags, User) {
        this.Topics = Topics
        this.Tags = Tags
        this.User = User
    }
}


const swmode = document.getElementById("swmode")
const localUser = localStorage.getItem("loged_user")?.toString()
let storageUser = new User("", "", "", "", "", "");

if (localUser) {
    console.log("auto connect")
    storageUser = JSON.parse(localUser)
    console.log(storageUser)

    switch (storageUser.theme) {
        case ("dark"):
            document.querySelector('body').setAttribute('data-theme', 'dark');
            console.log("default dark")
            swmode.checked = true
            break
        case ("light"):
            document.querySelector('body').setAttribute('data-theme', 'light');
            console.log("default light")
            swmode.checked = false
            break
    }

} else {
    console.log("to connect")
}





let publisher
let pp_publi
let list_tags = []
let list_topics = []
let list_cards = []
const display_tags = document.getElementById("display_tags")
const display_topics = document.getElementById("display_topics")
const style_mod = document.getElementById("style_mod")

async function fetch_all() {
    const tagsload = await fetch("http://localhost:8000/tags", {
        method: 'GET',
        headers: {
            "Accept": "application/json",
            "Content-type": "application/json; charset=UTF-8"
        }
    })
        .then((res) => {
            console.log(res)
            if (res.ok) {
                console.log("res.ok true")
                res.json().then(data => {
                    data.forEach(elt => {
                        list_tags.push(new Tags(elt.id_tags, elt.tags))
                    });
                    console.log(list_tags)

                    list_tags.forEach(elt => {
                        display_tags.innerHTML += `<h4 class="tag tags${elt.id_tags}" style= ""><span class="hover-underline-animation">${elt.tags}</span></h4>`
                        style_mod.innerHTML += `
.tags${elt.id_tags}::before{
    content: url(../../Assets/Images/icon_tag/tags${elt.id_tags}.svg);
    position: absolute;
    left: -20px;
}`
                    })

                    // list_pp.forEach(elt => {
                    //     document.getElementById(`imgpp${elt.id_pp}`).onclick = setpp
                    // })

                    // hiddenppvalue.value = document.getElementById(`imgpp${list_pp[0].id_pp}`).value
                    // showpp.style.backgroundImage = document.getElementById(`imgpp${list_pp[0].id_pp}`).style.backgroundImage


                })
            } else {
                console.log("res.ok false")
            }
        });



    const topicsload = await fetch("http://localhost:8000/topics", {
        method: 'GET',
        headers: {
            "Accept": "application/json",
            "Content-type": "application/json; charset=UTF-8"
        }
    })
        .then((res) => {
            // console.log(res)
            if (res.ok) {
                // console.log("res.ok true")
                res.json().then(data => {
                    data.forEach(elt => {

                        let actual_topic = new Topics(elt.id_topics, elt.titre, elt.description, elt.crea_date, elt.format_crea_date, elt.id_tags, elt.id_user)
                        list_topics.push(actual_topic)

                        const topicsload = fetch(`http://localhost:8000/users/${elt.id_user}`, {
                            method: 'GET',
                            headers: {
                                "Accept": "application/json",
                                "Content-type": "application/json; charset=UTF-8"
                            }
                        })
                            .then((res) => {
                                // console.log(res)
                                if (res.ok) {
                                    // console.log("res.ok true")
                                    res.json().then(data => {
                                        publisher = new User(data.id_user, data.pseudo, data.email, data.passwd, data.id_imagepp, data.theme)
                                        console.log(publisher)
                                        console.log(actual_topic)

                                        display_topics.innerHTML += `
                                            <div class="card">
                                                <div class="top_card">
                                                    <h4 class="user_card${publisher.id_imagepp}">${publisher.pseudo}</h4>
                                                    <p> &ensp; publié le ${actual_topic.format_crea_date}</p>
                                                </div>
                                                <div class="middle_card">
                                                    <h3 class="title_topic${actual_topic.id_tags}">${actual_topic.titre}</h3>
                                                </div>
                                                <div class="bottom_card">
                                                    <p>${actual_topic.description}</p>
                                                </div>
                                            </div>`

                                        style_mod.innerHTML += `
.title_topic${actual_topic.id_tags}::before {
    content: url(../../Assets/Images/icon_tag/tags${actual_topic.id_tags}.svg);
}
.title_topic${actual_topic.id_tags} {
    text-align: center;
}`

                                        const ppload = fetch(`http://localhost:8000/pp/${publisher.id_imagepp}`, {
                                            method: 'GET',
                                            headers: {
                                                "Accept": "application/json",
                                                "Content-type": "application/json; charset=UTF-8"
                                            }
                                        })
                                            .then((res) => {
                                                // console.log(res)
                                                if (res.ok) {
                                                    // console.log("res.ok true")
                                                    res.json().then(data => {
                                                        pp_publi = new Imagepp(data.id_pp, data.image_loc)

                                                        style_mod.innerHTML += `
.user_card${data.id_pp}::before {
    content: url(../../Assets/Images/profil/${pp_publi.image_loc});
}`


                                                    })
                                                } else {
                                                    console.log("res.ok false")
                                                }
                                            });

                                    })
                                } else {
                                    console.log("res.ok false")
                                }
                            });


                    });
                    console.log(list_topics)

                    //             list_topics.forEach(elt => {
                    //                 display_topics.innerHTML += `
                    //                 <div class="card">
                    //     <div class="top_card">
                    //         <h4 class="user_card">${elt.id_user}</h4>
                    //         <p>publié le ${elt.format_crea_date}</p>
                    //     </div>
                    //     <div class="middle_card">
                    //         <h3 class="title_topic">${elt.titre}</h3>
                    //     </div>
                    //     <div class="bottom_card">
                    //         <p>${elt.description}</p>
                    //     </div>
                    // </div>`

                    //             })

                })
            } else {
                console.log("res.ok false")
            }
        });
}



fetch_all()


















const openpopup = document.getElementById("openpopup")
const Popup = document.getElementById("favDialog")
const content = document.getElementById("contenu")
const close_popup = document.getElementById('close_pop')

const Createpost = document.getElementById('Createpost')
const open_create = document.getElementById("open_create")
const close_create = document.getElementById("close_create")




openpopup.onclick = function switch_theme() {
    Popup.showModal();
    content.style.position = 'fixed';
};
close_popup.addEventListener('click', function onClose() {
    Popup.close();
    content.style.position = 'initial'
});

open_create.onclick = function open_create() {
    Createpost.showModal();
};
// close_create.addEventListener('click', function onClose_create() {
//     Createpost.close();
// });




