<script lang='ts'>
  let { name, to, variant, size, ...rest } = $props()
</script>

<a class="button {variant} {size}" href={to} {...rest}>{name}</a>

<style>
  a {
    text-decoration: none;
    color: inherit;
  }

  .button {
    border-radius: var(--radius-small);
    padding: var(--spacing-3);

    cursor: pointer;
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