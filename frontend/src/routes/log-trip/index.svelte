<script lang="ts">
  let name = $state("")
  let world = $state("World")

  const handleSubmit = async (e: Event) => {
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

<form on:submit|preventDefault={handleSubmit}>
  <h1>Hello, {world}!</h1>

  <label>Name</label>
  <input type="text" bind:value={name} />

  <button type="submit">Log Trip</button>
</form>