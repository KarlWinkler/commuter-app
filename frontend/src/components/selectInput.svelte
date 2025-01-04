<script lang="ts">
  let { name, options, value = $bindable(1) }= $props()
</script>

<label for={name}>{name}</label>
<select name={name} bind:value>
  {#each options as option}
    <option value={option.key}>{option.value}</option>
  {/each}
</select>

<style>
  select {
    background-color: var(--white);
    padding: var(--spacing-3);
    border: var(--line-width) solid var(--black);
    border-radius: var(--radius-small);
  }

  option {
    background-color: var(--red); 
  }
</style>