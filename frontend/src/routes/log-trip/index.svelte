<script lang="ts">
  import Button from '/src/components/button.svelte'
  import NumberInput from '/src/components/numberInput.svelte'
  import SelectInput from '/src/components/selectInput.svelte'
  import TextInput from '/src/components/textInput.svelte'

  let name = $state("")
  let distance = $state("0")
  let mode = $state(1)
  let world = $state("World")

  const handleSubmit = async (e: Event) => {
    e.preventDefault()
    world = name
    const response = await fetch('http://localhost:3001/submit', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({"name": name}),
    })
  }
</script>

<form>
  <h1>Hello, {mode}!</h1>
  <div class="form">
    <TextInput name="Name" bind:value={name} />
    <NumberInput name="Distance" bind:value={distance} />
    <SelectInput name="Mode" bind:value={mode} options={[
      {key: 1, value: "Car"},
      {key: 2, value: "Bike"},
      {key: 3, value: "Bus"},
      {key: 4, value: "Train"}
    ]} />
  </div>
  <div class="submit">
    <Button onclick={handleSubmit} name="Log Trip" to="" variant="primary" />
  </div>
</form>

<style>
  form {
    display: flex;
    flex-direction: column;

    align-items: center;
  }

  .form {
    display: grid;
    grid-template-columns: 1fr auto;

    gap: var(--spacing-4);
    padding: var(--spacing-4)
  }

  .submit {
    width: fit-content;
  }
</style>