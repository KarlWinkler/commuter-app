<script>
  let { name, type, value = $bindable('') } = $props()
</script>

<label for={name}>{name}</label>
<input name={name} type="text" bind:value />

<style>
  input {
    padding: var(--spacing-3);
    border-radius: var(--radius-small);

    border: var(--line-width) solid var(--black);
  }
</style>