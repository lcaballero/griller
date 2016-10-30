<template>
<div class="message">{{ message }}</div>
</template>

<script>
export default {
  name: 'app',
  data () {
    return { message: "Hello, YO!" }
  }
}
</script>

<style>
.message {
  color: red;
}
</style>