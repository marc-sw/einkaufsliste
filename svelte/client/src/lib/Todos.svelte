<script>
  import { onMount } from 'svelte';

  let products = [];
  let name = "";
  let menge = 1;

  async function loadProducts() {
    const res = await fetch('/produkt');
    products = await res.json();
  }

  async function addProduct() {
    await fetch('/produkt', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name, menge, erledigt: false })
    });
    name = "";
    menge = 1;
    loadProducts();
  }

  async function toggleErledigt(product) {
    await fetch('/produkt', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        id: product.id,
        name: product.name,
        menge: product.menge,
        erledigt: !product.erledigt
      })
    });
    loadProducts();
  }

  async function updateProduct(product) {
    await fetch('/produkt', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(product)
    });
    loadProducts();
  }

  async function removeProduct(id) {
    await fetch(`/produkt/${id}`, { method: 'DELETE' });
    loadProducts();
  }

  onMount(loadProducts);
</script>

<style>
  .completed {
    text-decoration: line-through;
    opacity: 0.6;
  }
  input[type=number] {
    width: 60px;
  }
</style>

<h1>Products</h1>

<div>
  <input placeholder="Name" bind:value={name} />
  <input type="number" min="1" bind:value={menge} />
  <button on:click={addProduct}>Add</button>
</div>

<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Amount</th>
      <th>Actions</th>
    </tr>
  </thead>
  <tbody>
    {#each products as p}
      <tr>
        <td class={p.erledigt ? 'completed' : ''}>
          <input type="text" bind:value={p.name} on:change={() => updateProduct(p)} class={p.erledigt ? 'completed' : ''} />
        </td>
        <td>
          <input type="number" min="1" bind:value={p.menge} on:change={() => updateProduct(p)} />
        </td>
        <td>
          <button on:click={() => toggleErledigt(p)}>
            {p.erledigt ? 'Undo' : 'Complete'}
          </button>
          <button on:click={() => removeProduct(p.id)}>Delete</button>
        </td>
      </tr>
    {/each}
  </tbody>
</table>
