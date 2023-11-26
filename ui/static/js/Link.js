export default class Link {
  constructor(selector, location, notLink) {
    this.topicLinks = document.querySelectorAll(selector);
    this.location = location;
    this.notLink = notLink;
    this.addEventListeners();
  }

  addEventListeners() {
    this.topicLinks.forEach((link) => {
      link.addEventListener("click", (e) => {
        if (!e.target.classList.contains(this.notLink)) {
          this.linkClick(e);
        }
      });
    });
  }

  linkClick(event) {
    event.preventDefault();
    const topicId = event.currentTarget.getAttribute("data-id");
    window.location.href = `${this.location}${topicId}`;
  }
}
