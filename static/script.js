document.addEventListener('DOMContentLoaded', () => {
    const form = document.querySelector('form');
    const result = document.getElementById('result');

    form.addEventListener('submit', async (e) => {
        e.preventDefault();

        const inputs = [...form.querySelectorAll('input')].map(input => parseFloat(input.value));

        const apiEndpoints = {
            calculator1: '/api/calculator1',
            calculator2: '/api/calculator2',
            calculator3: '/api/calculator3',
        };

        const url = apiEndpoints[form.id];

        try {
            const response = await fetch(url, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ values: inputs })
            });

            const data = await response.json();
            console.log(data)
            result.textContent = `Результат: ${data.result}`;
        } catch (error) {
            result.textContent = 'Помилка.';
            console.error(error);
        }
    });
});