const openpopup = document.getElementById("openpopup")
const Popup = document.getElementById("favDialog")
const content = document.getElementById("contenu")
const close_popup = document.getElementById('close_pop')


openpopup.onclick = function switch_theme() {
    Popup.showModal();
    content.style.position = 'fixed';
};
close_popup.addEventListener('click', function onClose() {
    Popup.close();
    content.style.position = 'initial'
});
