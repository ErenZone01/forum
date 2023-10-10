function afficher(id) {
  let d2 = document.getElementById("com " + id);
  if (getComputedStyle(d2).display !== "none") {
    d2.style.display = "none";
    localStorage.setItem("element_caché_" + id, "caché"); // Stockez l'état caché
  } else {
    d2.style.display = "flex";
    localStorage.setItem("element_caché_" + id, "visible"); // Stockez l'état visible
  }
}

// Vérifiez l'état stocké lors du chargement de la page
document.addEventListener("DOMContentLoaded", function () {
  // Loop à travers les éléments
  let elements = document.querySelectorAll("[id^='com ']");
  elements.forEach(function (element) {
    let id = element.id.replace("com ", "");
    let etat = localStorage.getItem("element_caché_" + id);
    if (etat === "caché") {
      element.style.display = "none";
    } else if (etat === "visible") {
      element.style.display = "flex";
    }
  });
});

const textarea = document.querySelector("textarea");
textarea.addEventListener("keyup", (e) => {
  let scHeight = e.target.scrollHeight;
  textarea.style.height = "auto";
  textarea.style.height = `${scHeight}px`;
});

var fragment = window.location.hash;

// Supprimez le '#' du début de la chaîne si nécessaire
if (fragment.startsWith("#C")) {
  fragment = fragment.slice(1);
  var parts = fragment.split("%20");
  var partieSuivante = parts[1];
  console.log("Fragment de l'URL : " + partieSuivante);
  scrollToElementWithOffset(partieSuivante, 10)
}

// Fonction pour effectuer le défilement avec un décalage
function scrollToElementWithOffset(elementId, offset) {
  var element = document.getElementById(elementId);
  if (element) {
    var elementPosition = element.getBoundingClientRect().top + window.scrollY;
    window.scrollTo({
      top: elementPosition - offset,
      behavior: "smooth" // Pour un défilement fluide (peut être omis si vous ne le souhaitez pas)
    });
  }
}
