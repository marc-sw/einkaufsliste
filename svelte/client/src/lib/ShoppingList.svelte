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
    product.erledigt = !product.erledigt
    await fetch('/produkt', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        id: product.id,
        name: product.name,
        menge: product.menge,
        erledigt: product.erledigt
      })
    });
    products = products
    //loadProducts();
  }

  async function updateProduct(product) {
    await fetch('/produkt', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(product)
    });
    products = products
    //loadProducts();
  }

  async function removeProduct(product) {
    await fetch(`/produkt/${product.id}`, { method: 'DELETE' });
    //loadProducts();
    const index = products.indexOf(product)
    products.splice(index, 1)
    products = products
  }

  onMount(loadProducts);
</script>

<h1>Shopping List</h1>

<div class="spacing">
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
          <input type="number" min="1" bind:value={p.menge} on:change={() => updateProduct(p)} class={p.erledigt ? 'completed' : ''} />
        </td>
        <td>
          <button on:click={() => toggleErledigt(p)}>
            {p.erledigt ? 'Undo' : 'Complete'}
          </button>
          <button on:click={() => removeProduct(p)}>Delete</button>
        </td>
      </tr>
    {/each}
  </tbody>
</table>

<style>
  .completed {
    text-decoration: line-through;
    opacity: 0.6;
  }
  input[type=number] {
    width: 60px;
  }

  .spacing {
    margin-top: 1rem;
    margin-bottom: 1rem;
  }
</style>
