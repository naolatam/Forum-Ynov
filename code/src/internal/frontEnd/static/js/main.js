function toggleForm(isSignUp) {
    const container = document.getElementById("container");
    if (isSignUp) {
        container.classList.add("active");
    } else {
        container.classList.remove("active");
    }
}