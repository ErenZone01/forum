// const showHiddenInput = (passeordOverlayId, passwordInputId, toggleIconId) => {
//     const overlay = document.getElementById(passeordOverlayId),
//     input = document.getElementById(passwordInputId),
//     iconEye = document.getElementById(toggleIconId);

//     iconEye.addEventListener("click", () => {
//         if(input.type === "password"){
//             input.type = "text";
//             iconEye.classList.add("uil-show");
//         }else {
//             input.type = "password";
//             iconEye.classList.remove("uil-show")
//         }

//         overlay.classList.toggle("overlay-content");
//     });
// }

// showHiddenInput("password-overlay", "password-input", "toggle-visibility-icon");

// Sélectionnez les éléments HTML pertinents
const passwordInput = document.getElementById("password-input");
const toggleVisibilityIcon = document.querySelector(".toggle-visibility-icon");
const overlay = document.getElementById("password-overlay");

// Ajoutez un gestionnaire d'événements au clic sur l'icône de visibilité
toggleVisibilityIcon.addEventListener("click", () => {
    if (passwordInput.type === "password") {
        // Si le champ de mot de passe est masqué, le rendre visible
        passwordInput.type = "text";
    } else {
        // Sinon, le masquer
        passwordInput.type = "password";
    }
    overlay.classList.toggle("overlay-content");
});

// Sélectionnez les éléments HTML pertinents
const confirmationPasswordInput = document.getElementById("confirmation-password-input");
const toggleVisibilitySecondIcon = document.querySelector(".toggle-visibility-secondicon");
const confirmationOverlay = document.getElementById("confirmation-password-overlay");

// Ajoutez un gestionnaire d'événements au clic sur l'icône de visibilité
toggleVisibilitySecondIcon.addEventListener("click", () => {
    if (confirmationPasswordInput.type === "password") {
        // Si le champ de mot de passe est masqué, le rendre visible
        confirmationPasswordInput.type = "text";
    } else {
        // Sinon, le masquer
        confirmationPasswordInput.type = "password";
    }
    confirmationOverlay.classList.toggle("overlay-content");
});
