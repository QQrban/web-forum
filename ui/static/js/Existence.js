export default class Existence {
  constructor(selector, field, errorSelector) {
    this.selector = document.getElementById(selector);
    this.field = field;
    this.error = document.getElementById(errorSelector);
    this.isValid = false;

    this.checkInput = this.checkInput.bind(this);

    if (this.selector) {
      this.selector.addEventListener("change", this.checkInput);
    }
  }

  checkInput(event) {
    const inputValue = event.target.value;
    validate(this, inputValue);
  }
}

async function validate(my, value) {
  const field = my.field;
  try {
    const response = await fetch("/query/checkExistence", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        field: field,
        value: value,
      }),
    });
    const data = await response.json();
    if (data.ok) {
      my.isValid = true;
      my.error.innerHTML = "";
    } else {
      my.error.innerHTML = data.message;
      my.isValid = false;
    }
  } catch (error) {
    console.error("Error:", error);
    my.isValid = false;
  }
}
