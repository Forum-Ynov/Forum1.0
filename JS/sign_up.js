
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
                    showpp.style.backgroundImage = document.getElementById(`imgpp${list_pp[0].id_pp }`).style.backgroundImage


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










