<script>
  import { onMount } from 'svelte';

  let products = [];
  let name = "";
  let quantity = 1;

  async function loadProducts() {
    const res = await fetch('/shoppinglist-api/v1/products');
    products = await res.json();
  }

  async function addProduct() {
    await fetch('/shoppinglist-api/v1/products', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name, quantity, done: false })
    });
    name = "";
    quantity = 1;
    loadProducts();
  }

  async function toggleDone(product) {
    product.done = !product.done
    await fetch('/shoppinglist-api/v1/products', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        id: product.id,
        name: product.name,
        quantity: product.quantity,
        done: product.done
      })
    });
    products = products
  }

  async function updateProduct(product) {
    await fetch('/shoppinglist-api/v1/products', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(product)
    });
    products = products
  }

  async function removeProduct(product) {
    await fetch(`/shoppinglist-api/v1/products/${product.id}`, { method: 'DELETE' });
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
  <input type="number" min="1" bind:value={quantity} />
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
        <td class={p.done ? 'completed' : ''}>
          <input type="text" bind:value={p.name} on:change={() => updateProduct(p)} class={p.done ? 'completed' : ''} />
        </td>
        <td>
          <input type="number" min="1" bind:value={p.quantity} on:change={() => updateProduct(p)} class={p.done ? 'completed' : ''} />
        </td>
        <td>
          <button on:click={() => toggleDone(p)}>
            {p.done ? 'Undo' : 'Complete'}
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
