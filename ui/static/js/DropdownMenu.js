export default class DropdownMenu {
  constructor(selector, list, svg) {
    this.menu = document.querySelector(selector);
    this.list = document.querySelector(list);
    this.svg = document.querySelector(svg);
    this.toggleDropdown = this.toggleDropdown.bind(this);
    this.addEventListener();
  }

  addEventListener() {
    this.menu.addEventListener("click", this.toggleDropdown);

    document.addEventListener("click", (e) => {
      if (!this.menu.contains(e.target) && !this.list.contains(e.target)) {
        this.menu.classList.remove("active-topic");
        this.list.classList.remove("active-topic");
        this.svg.classList.remove("active-topic");
      }
    });
  }

  toggleDropdown() {
    this.menu.classList.toggle("active-topic");
    this.list.classList.toggle("active-topic");
    this.svg.classList.toggle("active-topic");
  }
}
