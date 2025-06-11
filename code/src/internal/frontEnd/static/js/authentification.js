function toggleForm(isSignUp) {
    const container = document.getElementById("container");
    if (isSignUp) {
        const url = new URL(window.location.href);
        url.searchParams.set("isRegister", "true");
        window.history.replaceState(null, null, url.toString());
        container.classList.add("active");
    } else {
        const url = new URL(window.location.href);
        url.searchParams.delete("isRegister");
        window.history.replaceState(null, null, url.toString());
        container.classList.remove("active");
    }
}

document.addEventListener("DOMContentLoaded", function () {
    const closeErrorBtn = document.getElementById('close-error');
    const errorPopUp = document.getElementById('error-popUp');
    closeErrorBtn.addEventListener('click', () => {
        errorPopUp.classList.add('hidden');
    }
    );
})