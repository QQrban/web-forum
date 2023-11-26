import PopUpModal from "./PopUpModal.js";
import DataFetcher from "./DataFetcher.js";
import { commentariesTemplate } from "./templates.js";
import { hideLoader, showLoader } from "./helpers.js";
import Like from "./Like.js";

export default class Comment extends DataFetcher {
  static endpoint = "/api/post/";
  static container = ".post-page__comment__container";
  static template = commentariesTemplate;

  constructor(logged) {
    super(Comment.endpoint, Comment.container, Comment.template);
    this.isLogged = logged;
    this.form = document.querySelector("#comment__body__add-comment");
    this.handleSubmit = this.handleSubmit.bind(this);
    this.addComment();
  }

  async handleSubmit(e) {
    e.preventDefault();
    const text = e.target[0].value;
    const location = window.location.pathname;
    const postID = location.substring(location.lastIndexOf("/") + 1);
    if (text.length === 0) {
      this.showErrorModal("Your comment can't be empty!");
      return;
    }
    try {
      showLoader();
      const response = await fetch("/create-new-comment", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          text: text,
          id: postID,
        }),
        credentials: "include",
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      } else {
        const data = await response.json();
        if (data.Ok) {
          this.fetchByID("ascending").then(() => {
            const successModal = new PopUpModal(
              "Success!",
              "Your comment has been added",
              "success",
              3500
            );
            successModal.showModal();
            e.target.reset();
            const reactions = new Like(
              ".likes-dislikes",
              "comment",
              "notActivity"
            );
          });
        } else {
          this.showErrorModal("Something went wrong :(");
        }
      }
    } catch (err) {
      this.showErrorModal("Something went wrong :(");
      console.error(err);
    } finally {
      hideLoader();
    }
  }

  showErrorModal(message) {
    const errorModal = new PopUpModal("Error!", message, "error", 3500);
    errorModal.showModal();
  }

  addComment() {
    if (!this.isLogged) {
      return;
    }
    this.form.addEventListener("submit", this.handleSubmit);
  }

  removeComment() {
    this.form.removeEventListener("submit", this.handleSubmit);
  }
}
