const mainAudioEl = document.createElement("audio");
document.body.appendChild(mainAudioEl);

document.querySelector("#btn-play").addEventListener("click", () => {
    document.body.removeChild(document.querySelector("#start"));
    document.querySelector("#app").classList.remove("hidden");
    randonlyPlayMainSong();
});

async function randonlyPlayMainSong() {
    const songsFetch = await fetch("/public/songs/available_songs.json");
    /** @type string[] */
    const availableSongs = await songsFetch.json();
    const pickedSongIdx = Math.round(
        Math.random() * (availableSongs.length - 1),
    );

    mainAudioEl.src = `/public/songs/${availableSongs[pickedSongIdx]}`;
    mainAudioEl.play();
}
