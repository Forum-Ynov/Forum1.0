
class user {
    constructor(id_user, pseudo, email, passwd, id_imagepp) {
        this.id_user = id_user
        this.pseudo = pseudo
        this.email = email
        this.passwd = passwd
        this.id_imagepp = id_imagepp
    }
}


let message = document.getElementById("message");
let signup = document.getElementById("signup");
let in_pseudo = document.getElementById("in_pseudo");
let in_passwd = document.getElementById("in_passwd");
let form_log = document.getElementById("form_log");

let actual_user = new user("", "", "", "", "");

const fetchUsers = async (name) => {
    fetch(`http://localhost:8000/userpseudo/${name}`)
        .then((res) => {
            if (res.ok) {
                res.json().then(data => {
                    actual_user.id_user = data.id_user;
                    actual_user.pseudo = data.pseudo;
                    actual_user.email = data.email;
                    actual_user.passwd = data.passwd;
                    actual_user.id_imagepp = data.id_imagepp;
                })
                // message.innerHTML = ""
            } else {
                console.log("ERREUR");
                message.innerHTML = "User not found";
            }
        });
};


const usersDisplay = async (name) => {
    await fetchUsers(name);
};

form_log.addEventListener("submit", function (event) {
    event.preventDefault();
    let names = in_pseudo.value;
    let pass = in_passwd.value;
    usersDisplay(names)
    console.log(actual_user.passwd, pass)
    // if(actual_user.passwd !== undefined){
    if (actual_user.passwd == pass) {
        alert("connected")
        // this.submit();
    } else {
        message.innerHTML = "password incorect"
        event.preventDefault();
    }
    // }

});
