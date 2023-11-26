import PopUpModal from "./PopUpModal.js";
import { showLoader, hideLoader } from "./helpers.js";

export default class Like {
  constructor(selector, endpoint, location) {
    this.selector = document.querySelectorAll(selector);
    this.endpoint = endpoint;
    this.location = location;
    this.handleLike = this.handleLike.bind(this);
    this.modal = document.querySelector("#remove-like-modal");

    this.addLike();
  }

  async handleLike(e) {
    const currentTarget = e.currentTarget.closest(
      "[data-id][data-likes][data-dislikes]"
    );
    const id = currentTarget.getAttribute("data-id");

    try {
      const reaction = this.getValue(e);
      if (this.location === "notActivity") {
        const response = await fetch(
          `/api/reaction/${this.endpoint}/${id}/${reaction}`
        );

        if (!response.ok) {
          if (response.status === 401) {
            const errorMessage = new PopUpModal(
              "Error!",
              "To like or dislike, please log in or sign up",
              "error",
              2500
            );
            errorMessage.showModal();
          }
        } else {
          const data = await response.json();
          if (data.Ok) {
            let dataLikes = currentTarget.getAttribute("data-likes");
            let dataDislikes = currentTarget.getAttribute("data-dislikes");
            let like = currentTarget.querySelector("#likes");
            let dislike = currentTarget.querySelector("#dislikes");
            let likeNumber;
            let dislikeNumber;

            if (dataLikes !== null && dataDislikes !== null) {
              if (data.Add === "like") {
                likeNumber = +like.innerHTML.trim() + 1;
                currentTarget.setAttribute(
                  "data-likes",
                  parseInt(dataLikes) + 1
                );
                like.innerHTML = likeNumber;
              } else if (data.Delete === "like") {
                likeNumber = +like.innerHTML.trim() - 1;
                currentTarget.setAttribute(
                  "data-likes",
                  parseInt(dataLikes) - 1
                );
                like.innerHTML = likeNumber;
              } else if (data.UpdateTo === "like") {
                likeNumber = +like.innerHTML.trim() + 1;
                dislikeNumber = +dislike.innerHTML.trim() - 1;
                currentTarget.setAttribute(
                  "data-likes",
                  parseInt(dataLikes) + 1
                );
                currentTarget.setAttribute(
                  "data-dislikes",
                  parseInt(dataDislikes) - 1
                );
                like.innerHTML = likeNumber;
                dislike.innerHTML = dislikeNumber;
              }
              if (data.Add === "dislike") {
                dislikeNumber = +dislike.innerHTML.trim() + 1;
                currentTarget.setAttribute(
                  "data-dislikes",
                  parseInt(dataDislikes) + 1
                );
                dislike.innerHTML = dislikeNumber;
              } else if (data.Delete === "dislike") {
                dislikeNumber = +dislike.innerHTML.trim() - 1;
                currentTarget.setAttribute(
                  "data-dislikes",
                  parseInt(dataDislikes) - 1
                );
                dislike.innerHTML = dislikeNumber;
              } else if (data.UpdateTo === "dislike") {
                dislikeNumber = +dislike.innerHTML.trim() + 1;
                likeNumber = +like.innerHTML.trim() - 1;
                currentTarget.setAttribute(
                  "data-likes",
                  parseInt(dataLikes) - 1
                );
                currentTarget.setAttribute(
                  "data-dislikes",
                  parseInt(dataDislikes) + 1
                );
                like.innerHTML = likeNumber;
                dislike.innerHTML = dislikeNumber;
              }
            }
          }
        }
      } else {
        const deleteText = this.modal.querySelector(".remove-like__heading");
        if (reaction === "dislike") {
          deleteText.textContent =
            "Are you sure you want to dislike this post?";
        }
        if (reaction === "like") {
          deleteText.textContent = "Are you sure you want to remove your like?";
        }
        this.setupModal(id, reaction);
      }
    } catch (error) {
      console.error(error);
    }
  }

  setupModal(id, reaction) {
    this.currentId = id;
    this.currentReaction = reaction;
    this.modal.classList.remove("hidden");
    this.closeModalBtn = document.querySelector("#remove-like-modal__close");
    if (this.closeModalBtn) {
      this.closeModalBtn.addEventListener("click", () => this.closeModal());
    }

    const yesBtn = this.modal.querySelector(".btn__yes");
    const noBtn = this.modal.querySelector(".btn__no");
    if (yesBtn && noBtn) {
      yesBtn.removeEventListener("click", this.yesBtnClickHandler);
      noBtn.removeEventListener("click", this.noBtnClickHandler);

      this.yesBtnClickHandler = () => this.confirmDecision();
      this.noBtnClickHandler = () => {
        this.closeModal();
      };

      yesBtn.addEventListener("click", this.yesBtnClickHandler);
      noBtn.addEventListener("click", this.noBtnClickHandler);
    }
  }

  async confirmDecision() {
    try {
      const response = await fetch(
        `/api/reaction/${this.endpoint}/${this.currentId}/${this.currentReaction}`
      );
      if (!response.ok) {
        const errorMessage = new PopUpModal(
          "Error!",
          "Something went wrong :(",
          "error",
          2500
        );
        errorMessage.showModal();
      } else {
        showLoader();
        const data = await response.json();
        if (data) {
          setTimeout(() => {
            window.location.reload();
          }, 100);
        }
      }
    } catch (error) {
      console.log("Error deleting post: ", error);
    }
  }

  closeModal() {
    this.modal.classList.add("hidden");
  }

  getValue(e) {
    if (
      e.target.classList.contains("comments__likes") ||
      e.target.parentNode.classList.contains("comments__likes")
    ) {
      return "like";
    } else if (
      e.target.classList.contains("comments__dislikes") ||
      e.target.parentNode.classList.contains("comments__dislikes")
    ) {
      return "dislike";
    }
  }

  addLike() {
    this.selector.forEach((item) =>
      item.addEventListener("click", this.handleLike)
    );
  }
}
