
import {User} from "./user_class.js"

let message = document.getElementById("message");
let in_pseudo = document.getElementById("in_pseudo");
let in_passwd = document.getElementById("in_passwd");
let form_log = document.getElementById("form_log");

let actual_user = new User("", "", "", "", "", "");
let default_user = new User("", "", "", "", "", "");

form_log.addEventListener("submit", async function (event) {
    event.preventDefault();
    let names = in_pseudo.value;
    let pass = in_passwd.value;

    if (storageUser.id_user == default_user.id_user && storageUser.pseudo == default_user.pseudo && storageUser.email == default_user.email && storageUser.passwd == default_user.passwd && storageUser.id_imagepp == default_user.id_imagepp && storageUser.theme == default_user.theme) {

        const r = await fetch("http://localhost:8000/apiForum/login", {
            method: 'POST',
            headers: {
                "Accept": "application/json",
                "Content-type": "application/json; charset=UTF-8"
            },
            body: JSON.stringify({ pseudo: names, passwd: pass })
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
                        alert("connected")
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
    }
});

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