<script setup lang="ts"></script>

<template>
  <div class="flex h-full">
    <nav class="nav top">Navigator</nav>
    <section class="flex-grow nav scrollable-y">
      <p v-for="n in 1">View {{ n }}</p>
    </section>
    <footer class="nav bottom">footer</footer>
  </div>
</template>

<style lang="css" scoped>
.nav {
  padding: 0.75rem 0.5rem;
}

.nav.top {
  border-bottom: 1px solid black;
}

.nav.bottom {
  border-top: 1px solid black;
}

.scrollable-y {
  overflow-y: scroll;
}
</style>
