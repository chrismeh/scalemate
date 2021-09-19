function updateImageSRC() {
    const root = encodeURIComponent(document.getElementById("root").value);
    const scale = encodeURIComponent(document.getElementById("scale").value);
    const tuning = encodeURIComponent(document.getElementById("tuning").value);

    document.getElementById("scale-image").src = "/scale?root=" + root + "&type=" + scale + "&tuning=" + tuning;
}