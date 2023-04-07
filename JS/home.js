
import { User } from "./user_class.js";
import { Imagepp } from "./pp_class.js";
import { Tags } from "./tags_class.js";
import { Topics } from "./topics_class.js";
import { Messages } from "./messages_class.js";

let list_tags = []
const display_tags = document.getElementById("display_tags")
const style_mod = document.getElementById("style_mod")

async function loadpp() {
    const r = await fetch("http://localhost:8000/tags", {
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
}

loadpp()
















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
close_create.addEventListener('click', function onClose_create() {
    Createpost.close();
});




