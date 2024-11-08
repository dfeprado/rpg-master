<script setup lang="ts">
import Master from './ui/master/Master.vue'
import Player from './ui/player/Player.vue'

const cookies = document.cookie.split(';').map((e) => e.trim())
const isMaster = cookies.length > 0 && cookies.find((e) => e == 'UI=MASTER')

const viewComponent = isMaster ? Master : Player
</script>

<template>
  <component :is="viewComponent"></component>
</template>
