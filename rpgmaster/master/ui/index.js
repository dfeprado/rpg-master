const mainAudioEl = document.createElement("audio");
document.body.appendChild(mainAudioEl);
const btnPlay = document.querySelector("#btn-play");

// TODO remove after dev
window.addEventListener("DOMContentLoaded", () => btnPlay.click());

btnPlay.addEventListener("click", () => {
    document.body.removeChild(document.querySelector("#start"));
    document.querySelector("#app").classList.remove("hidden");
    // TODO uncomment after dev
    // randonlyPlayMainSong();
});

async function randonlyPlayMainSong() {
    const songsFetch = await fetch("/static/songs/available_songs.json");
    /** @type string[] */
    const availableSongs = await songsFetch.json();
    const pickedSongIdx = Math.round(
        Math.random() * (availableSongs.length - 1),
    );

    mainAudioEl.src = `/static/songs/${availableSongs[pickedSongIdx]}`;
    mainAudioEl.play();
}
