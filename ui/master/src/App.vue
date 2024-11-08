<script setup lang="ts">
import { ref } from 'vue'
const msg: string = 'Hello, World!'
const response = ref<{ hello: string }>()

fetch('http://localhost:8080/api/')
  .then((r) => r.json())
  .then((v) => (response.value = v))
  .catch((err) => (response.value = { hello: err }))
</script>

<template>
  <h1>{{ msg }}</h1>
  <p v-if="response">
    {{ response.hello }}
  </p>
</template>
