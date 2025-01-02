<script lang="ts">
  let name = $state("")
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

<form onsubmit={handleSubmit}>
  <h1>Hello, {world}!</h1>

  <label>
    Name
    <input type="text" bind:value={name} />
  </label>
  <button type="submit">Log Trip</button>
</form>