<script>
  let { name, to, variant, size, ...rest } = $props()
</script>

<div class="button {variant} {size}" {...rest}>
  <a href={to}>{name}</a>
</div>

<style>
  a {
    text-decoration: none;
    color: inherit;
  }

  .button {
    border-radius: var(--radius-small);
    padding: var(--spacing-3);
  }

  .button:hover {
    background-color: var(--light-red);
    color: var(--red);
  }

  .primary {
    border: var(--line-width) solid var(--red);
    border-radius: var(--radius-small);

    background-color: var(--red);
    color: var(--white);
  }

  .secondary {
    border: var(--line-width) solid var(--red);
    border-radius: var(--radius-small);

    background-color: var(--white);
    color: var(--red);
  }

  .tertiary {
    border: var(--line-width) solid var(--black);
    border-radius: var(--radius-small);
  }

  .small {
    font-size: var(--font-size-small);
  }

  .large {
    font-size: var(--font-size-large);
    border-radius: var(--radius-large)
  }

</style>