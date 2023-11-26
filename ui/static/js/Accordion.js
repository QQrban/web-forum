export default class Accordion {
  constructor(selector) {
    this.accordions = Array.from(document.getElementsByClassName(selector));
    this.initializeAccordions();
  }

  initializeAccordions() {
    this.accordions.forEach((acc) => {
      acc.classList.add("active");
      const panel = acc.nextElementSibling;
      panel.style.display = "block";

      acc.addEventListener("click", () => {
        acc.classList.toggle("active");
        panel.style.display =
          panel.style.display === "block" ? "none" : "block";
      });
    });
  }
}
