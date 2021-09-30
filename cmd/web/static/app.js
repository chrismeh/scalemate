function updateImageSRC() {
    const root = encodeURIComponent(document.getElementById("root").value);
    const scale = encodeURIComponent(document.getElementById("scale").value);
    const tuning = encodeURIComponent(document.getElementById("tuning").value);
    const url = "/api/scale?root=" + root + "&type=" + scale + "&tuning=" + tuning;

    fetch(url)
        .then(resp => resp.json())
        .then(json => {
            console.log(json)
            document.getElementById("scale-image").src = `data:image/png;base64,${json.picture}`;
        })
}