<script lang="ts">
  import Button from "/src/components/button.svelte"
</script>

<div class="header-wrapper">
  <div class="header-links">
    <Button name="Home" to="/" />
    <Button name="Log Trip" to="/log-trip" />
  </div>
</div>

<style>
  .header-wrapper {
    display: flex;

    background: white;
    padding: var(--spacing-4);
  }

  .header-links {
    display: flex;
  }
</style>