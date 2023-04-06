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


