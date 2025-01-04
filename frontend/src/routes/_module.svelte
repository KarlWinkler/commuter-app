<script lang="ts">
  import Header from "./header.svelte"
  import Footer from "./footer.svelte"
</script>

<div class="app">
  <Header />
  <div class="content-wrapper">
    <slot />
  </div>
  <Footer />
</div>

<style>
  .app {
    min-height: 100vh;
    display: grid;
    grid-template-rows: auto 1fr auto;
    color: var(--black);
    padding: 0;
  }
  .content-wrapper {
    display: flex;

    flex-direction: column;
    justify-content: center;
    align-items: center;
    
    padding: var(--spacing-7);
  }
</style>