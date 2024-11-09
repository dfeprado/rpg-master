<script setup lang="ts">
import Master from './ui/master/Master.vue'
import Player from './ui/player/Player.vue'

const isMaster = window.location.hostname == 'localhost' || window.location.hostname == '127.0.0.1'
const viewComponent = isMaster ? Master : Player
</script>

<template>
  <component :is="viewComponent"></component>
</template>
