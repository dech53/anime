<template>
  <div id="app">
    <video id="video-player" controls autoplay></video>
    <div v-if="animes.length === 0">获取信息失败</div>
    <div v-for="anime in animes" :key="anime.Name" class="anime-item">
      <img
        :src="anime.Cover"
        :alt="anime.Name"
        class="cover"
        @click="selectAnime(anime)"
      />
      <div>
        <h2 @click="selectAnime(anime)">{{ anime.Name }}</h2>
        <p>Season: {{ anime.Season }}</p>
        <p>
          Episodes:
          <br>
          <br>
          <span
            v-for="episode in anime.EpisodeCount"
            :key="episode"
            class="episode-number"
            @click="playEpisode(anime, episode)"
          >
            {{ episode }}
          </span>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Hls from "hls.js";

export default {
  data() {
    return {
      animes: [],
      currentVideo: null,
      hls: null,
    };
  },
  methods: {
    async fetchAnimeInfo() {
      try {
        const response = await axios.get("http://192.168.1.4:3421/getInfos");
        this.animes = response.data.result; // 这里使用 .result
      } catch (error) {
        console.error("Failed to fetch anime info:", error);
      }
    },
    selectAnime(anime) {
      this.currentVideo = null;
      this.selectedAnime = anime;
    },
    playEpisode(anime, episode) {
      const videoElement = document.getElementById("video-player");
      const episodePath = `/hls/${anime.Year}/${anime.Name}/${anime.Season}/${episode}/output.m3u8`;
      const videoUrl = `http://192.168.1.4:3421${episodePath}`;

      // 清理之前的 HLS 实例
      if (this.hls) {
        this.hls.destroy();
      }

      // 初始化 HLS.js
      if (Hls.isSupported()) {
        this.hls = new Hls();
        this.hls.loadSource(videoUrl);
        this.hls.attachMedia(videoElement);
        this.hls.on(Hls.Events.MANIFEST_PARSED, function () {
          videoElement.play();
        });
      } else if (videoElement.canPlayType("application/vnd.apple.mpegurl")) {
        // 对于支持 HLS 的浏览器（例如 Safari）
        videoElement.src = videoUrl;
        videoElement.addEventListener("loadedmetadata", function () {
          videoElement.play();
        });
      }
    },
  },
  mounted() {
    this.fetchAnimeInfo();
  },
};
</script>

<style>
body {
  font-family: Arial, sans-serif;
  background-color: #f4f4f4;
  margin: 0;
  padding: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

h1 {
  color: #333;
  margin-bottom: 20px;
  text-align: center;
}

.anime-item {
  display: flex;
  align-items: center;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
  padding: 15px;
  width: 80%;
  transition: transform 0.2s;
}

.anime-item:hover {
  transform: scale(1.02);
}

.cover {
  width: 120px;
  height: 180px;
  margin-right: 20px;
  cursor: pointer;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

h2 {
  cursor: pointer;
  margin: 0;
  font-size: 24px;
  color: #333;
}

p {
  margin: 5px 0;
  color: #666;
}

.episode-number {
  margin-right: 10px;
  cursor: pointer;
  background-color: #007bff;
  color: #fff;
  padding: 5px 10px;
  border-radius: 4px;
  transition: background-color 0.2s;
}

.episode-number:hover {
  background-color: #0056b3;
}

video {
  width: 60%;
  height: auto;
  border-radius: 8px;
  margin-top: 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

@media (max-width: 768px) {
  .anime-item {
    flex-direction: column;
    align-items: flex-start;
  }

  .cover {
    width: 100px;
    height: 150px;
    margin-bottom: 10px;
  }

  video {
    width: 100%;
  }
}
</style>
