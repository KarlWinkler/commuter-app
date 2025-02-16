<script lang="ts">
  import Button from '/src/components/button.svelte'
  import NumberInput from '/src/components/numberInput.svelte'
  import SelectInput from '/src/components/selectInput.svelte'
  import TextInput from '/src/components/textInput.svelte'
  import { goto } from "@roxi/routify";

  type Mode = {
    ID: number
    Name: {
      String: string
    }
  }

  let name = $state("")
  let distance = $state(0)
  let mode = $state(1)

  let redirect = $goto

  let modes = $state([])
  $effect(() => async () => {
    const response = await fetch('http://localhost:3001/modes')
    const data = await response.json()
    modes = data.Result.map((mode: Mode) => ({key: mode.ID, value: mode.Name.String}))
  })

  const handleSubmit = async (e: Event) => {
    e.preventDefault()
    let body = {
      "name": name,
      "distance": distance,
      "mode": mode,
    }
    const response = await fetch('http://localhost:3001/submit', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    })

    console.log(response)
    redirect('/success')
  }
</script>

<form>
  <h1>Log Your Trip!</h1>
  <div class="form">
    <TextInput name="Name" bind:value={name} />
    <NumberInput name="Distance" bind:value={distance} />
    <SelectInput name="Mode" bind:value={mode} options={modes} />
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
    margin-top: var(--spacing-4);
    width: fit-content;
  }
</style>