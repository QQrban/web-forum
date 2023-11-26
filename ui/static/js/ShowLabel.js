export default class ShowLabel {
  constructor(buttonSelector, labelSelector, timer) {
    this.button = document.querySelector(buttonSelector);
    this.label = document.querySelector(labelSelector);
    this.timeOut = null;
    this.timer = timer;
    this.addEventListeners();
  }

  addEventListeners() {
    this.button.addEventListener("mouseenter", () => this.onMouseEnter());
    this.button.addEventListener("mouseleave", () => this.onMouseOut());
  }

  onMouseEnter() {
    this.timeOut = setTimeout(() => {
      this.label.style.display = "block";
    }, this.timer);
  }

  onMouseOut() {
    this.label.style.display = "none";
    clearTimeout(this.timeOut);
  }

  destroy() {
    if (this.button) {
      this.button.removeEventListener("mouseenter", this.onMouseEnter);
      this.button.removeEventListener("mouseleave", this.onMouseOut);
    }

    if (this.timeOut) {
      clearTimeout(this.timeOut);
    }

    this.label.style.display = "none";
  }
}
