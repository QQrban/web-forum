import PopUpModal from "./PopUpModal.js";
import DataFetcher from "./DataFetcher.js";
import { postsTemplate } from "./templates.js";
import Link from "./Link.js";
import { hideLoader, showLoader } from "./helpers.js";
import Like from "./Like.js";
import Filters from "./Filters.js";

export default class CreatePost {
  constructor(logged) {
    this.formSelector = document.querySelector("#new-post__modal__form");
    this.button = document.querySelector("#topic__header__btn-svg");
    this.avatarModal = document.querySelector("#select-avatar");
    this.modal = document.querySelector(".new-post__modal");
    this.overlay = document.querySelector("#overlay");
    this.closeButton = document.querySelector("#new-post__modal__close-svg");
    this.isLogged = logged;
    this.isErrorModalActive = false;
    this.showModalBound = this.showModal.bind(this);
    this.hideModalBound = this.hideModal.bind(this);
    this.showErrorBound = this.showError.bind(this);
    this.documentClickHandlerBound = this.documentClickHandler.bind(this);

    this.updateEventListeners();
  }

  updateEventListeners() {
    this.removeEventListeners();
    if (this.isLogged) {
      this.initLoggedInEventListeners();
    } else {
      this.removeEventListeners();
      this.initLoggedOutEventListeners();
    }
  }

  initLoggedInEventListeners() {
    this.button.addEventListener("click", this.showModalBound);
    this.closeButton.addEventListener("click", this.hideModalBound);
    document.addEventListener("click", this.documentClickHandlerBound);
    this.validateAndSubmit();
  }

  initLoggedOutEventListeners() {
    this.button.addEventListener("click", () => {
      if (!this.isErrorModalActive) {
        this.showErrorBound();
      }
    });
  }

  documentClickHandler(e) {
    if (
      !this.modal.contains(e.target) &&
      !this.button.contains(e.target) &&
      !this.avatarModal.contains(e.target)
    ) {
      this.hideModal();
    }
  }

  removeEventListeners() {
    this.button.removeEventListener("click", this.showModalBound);
    this.button.removeEventListener("click", this.showErrorBound);
    this.closeButton.removeEventListener("click", this.hideModalBound);
    document.removeEventListener("click", this.documentClickHandlerBound);
  }

  showModal() {
    this.modal.style.display = "block";
    this.overlay.style.display = "block";
  }

  hideModal() {
    this.modal.style.display = "none";
    this.overlay.style.display = "none";
  }

  showError() {
    if (this.isLogged) {
      return;
    }
    this.isErrorModalActive = true;
    const errorPopup = new PopUpModal(
      "Error",
      "You need to be logged in to create posts.",
      "error",
      3000
    );
    errorPopup.showModal();
    setTimeout(() => {
      this.isErrorModalActive = false;
    }, 3000);
  }

  validateAndSubmit() {
    this.formSelector.addEventListener("submit", async (e) => {
      const errorMsg = document.querySelector("#add__post__error");
      const location = window.location.pathname;
      const topicID = location.substring(location.lastIndexOf("/") + 1);
      e.preventDefault();
      const title = e.target[0].value;
      const description = e.target[1].value;
      const tagsFromInput = e.target[2].value;
      const tagsRegex = /^[A-Za-z\s,]+$/;
      if (title === "" || description === "" || tagsFromInput === "") {
        errorMsg.innerHTML = "All fields must be filled out.";
        return;
      }
      if (!tagsRegex.test(tagsFromInput)) {
        errorMsg.innerHTML =
          "Each tag must be between 3 and 15 characters long and should be divided by comma";
        return;
      }
      let tags = [];
      let tagString = "";
      if (tagsFromInput.length != 0) {
        tagString = tagsFromInput.replace(/\s/g, "");
        tags = tagsFromInput.split(",").map((tag) => tag.trim());
        for (let i = 0; i < tags.length; i++) {
          if (tags[i].length < 3 || tags[i].length > 15) {
            errorMsg.innerHTML =
              "Each tag must be between 3 and 15 characters long.";
            return;
          }
        }
      }
      if (title.length < 5 || title.length > 50) {
        errorMsg.innerHTML = "Title must be between 5 and 50 characters long.";
        return;
      } else if (description.length < 50 || description.length > 1500) {
        errorMsg.innerHTML =
          "Description must be between 50 and 1500 characters long.";
        return;
      } else if (tags.length < 1 || tags.length > 10) {
        errorMsg.innerHTML = "Tags can be a minimum of 1 and a maximum of 10.";
        return;
      } else {
        errorMsg.innerHTML = "";
      }

      const creationTime = new Date().toISOString();
      showLoader();
      try {
        const response = await fetch("/create-new-post", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            title: title,
            content: description,
            tags: tagString,
            categoryId: topicID,
            created: creationTime,
          }),
          credentials: "include",
        });
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        } else {
          const data = await response.json();
          if (data.Ok) {
            const selectedTopic = new DataFetcher(
              "/api/topic/",
              ".all-posts__container",
              postsTemplate
            );
            selectedTopic.fetchByID("descending").then(() => {
              const topicLinks = new Link(
                ".all-posts__item",
                "/post/",
                "not__link"
              );
              const reactions = new Like(".likes-dislikes", "post", "notActivity");
              const sortByDate = new Filters();
              sortByDate.addTags();
            });
            e.target.reset();
            this.hideModal();
            const successModal = new PopUpModal(
              "Success!",
              "Your post has been added!",
              "success",
              3500
            );
            successModal.showModal();
          } else {
            errorMsg.innerHTML = data.Message;
          }
        }
      } catch (error) {
        console.error(error);
      } finally {
        hideLoader();
      }
    });
  }
}
