
import { User } from "./user_class.js"
import { Imagepp } from "./pp_class.js"

let list_pp = []

async function loadpp() {
    const r = await fetch("http://localhost:8000/pp", {
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
                        list_pp.push(new Imagepp(elt.id_pp, elt.image_loc))
                    });
                    console.log(list_pp)

                    list_pp.forEach(elt => {
                        ppcontent.innerHTML += `<option id="imgpp${elt.id_pp}" value="${elt.id_pp}" onclick = setpp()
                        style="background: no-repeat; background-image:url(../../Assets/Images/profil/${elt.image_loc}); height: 64px; width: 64px;">
                    </option >`

                    })

                    list_pp.forEach(elt => {
                        document.getElementById(`imgpp${elt.id_pp}`).onclick = setpp
                    })

                    hiddenppvalue.value = document.getElementById(`imgpp${list_pp[0].id_pp}`).value
                    showpp.style.backgroundImage = document.getElementById(`imgpp${list_pp[0].id_pp}`).style.backgroundImage


                })
            } else {
                console.log("res.ok false")
            }
        });
}

loadpp()


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


let passtotext = document.getElementById("passtotext")
let imgpass = document.getElementById("imgpass")

passtotext.onclick = function () {
    if (in_passwd.type == "password") {
        in_passwd.type = "text"
        imgpass.src = "../../Assets/Images/icon_input/eye-off.svg"
    } else {
        in_passwd.type = "password"
        imgpass.src = "../../Assets/Images/icon_input/eye.svg"
    }

}

const showpp = document.getElementById("showpp")
const ppcontent = document.getElementById("ppcontent")
const hiddenppvalue = document.getElementById("hiddenppvalue")

showpp.onclick = toggleselector

function toggleselector() {
    if (ppcontent.style.display != "none") {
        ppcontent.style.display = "none"
    } else {
        ppcontent.style.display = "flex"
    }

}


function setpp() {
    toggleselector()
    hiddenppvalue.value = this.value
    showpp.style.backgroundImage = this.style.backgroundImage
}


let message = document.getElementById("message");
let in_pseudo = document.getElementById("in_pseudo");
let in_email = document.getElementById("in_email");
let in_passwd = document.getElementById("in_passwd");
let ppvalue = document.getElementById("hiddenppvalue")
let form_sign = document.getElementById("form_sign");
const body = document.querySelector('body');

let actual_user = new User("", "", "", "", "", "");
let default_user = new User("", "", "", "", "", "");

form_sign.addEventListener("submit", async function (event) {
    event.preventDefault();
    let names = in_pseudo.value;
    let emailed = in_email.value;
    let pass = in_passwd.value;
    let imagepp = ppvalue.value;
    let user_theme = body.getAttribute('data-theme');


    const r = await fetch("http://localhost:8000/adduser", {
        method: 'POST',
        headers: {
            "Accept": "application/json",
            "Content-type": "application/json; charset=UTF-8"
        },
        body: JSON.stringify({ pseudo: names, email: emailed, passwd: pass, id_imagepp: parseInt(imagepp), theme: user_theme })
    })
        .then((res) => {
            console.log(res)
            if (res.ok) {
                console.log("res.ok true")
                res.json().then(data => {
                    actual_user.id_user = data.id_user;
                    actual_user.pseudo = data.pseudo;
                    actual_user.email = data.email;
                    actual_user.passwd = data.passwd;
                    actual_user.id_imagepp = data.id_imagepp;
                    actual_user.theme = data.theme
                    console.log(actual_user)
                    alert("created")
                    localStorage.setItem("loged_user", JSON.stringify(actual_user))
                })
            } else {
                console.log("res.ok false")
                res.json().then(data => {
                    console.log("ERREUR");
                    message.innerHTML = data.message;
                    actual_user.id_user = ""
                    actual_user.pseudo = ""
                    actual_user.email = ""
                    actual_user.passwd = ""
                    actual_user.id_imagepp = ""
                    actual_user.theme = ""
                })
            }
        });
});






