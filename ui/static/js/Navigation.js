export default class Navigation {
  constructor(buttonSelectors) {
    this.buttons = buttonSelectors.map((selector) =>
      document.querySelector(selector)
    );
    this.addEventListeners();
  }

  addEventListeners() {
    this.buttons.forEach((btn) => btn.addEventListener("click", this.toHome));
  }

  toHome() {
    window.location.href = "/";
  }
}
