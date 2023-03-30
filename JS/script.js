const toggleMenu = () => document.body.classList.toggle("open");

const body = document.querySelector('body');
// const heart = document.getElementById('like');
// const heart_red = document.getElementById('like_red')
// console.log(heart);
function switch_theme() {
    if (test.checked) {
        console.log("dark");
        body.setAttribute('data-theme', 'dark');
    }
    else {
        console.log("light");
        body.setAttribute('data-theme', 'light');
    }
}

// heart.addEventListener("click", function (e) {
//     console.log(e);
//     heart.style.display = 'none';
//     heart_red.style.opacity = 1
// });

// heart_red.addEventListener("click", function (e) {
//     heart_red.style.opacity = '1'
// })

