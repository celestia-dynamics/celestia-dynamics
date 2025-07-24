document.getElementById('simulate-btn').addEventListener('click', function() {
    const angle = parseFloat(document.getElementById('angle').value);
    const velocity = parseFloat(document.getElementById('velocity').value);
    const airResistance = parseFloat(document.getElementById('air-resistance').value);
    const windSpeed = parseFloat(document.getElementById('wind-speed').value);

    const { trajectory, range, timeOfFlight, finalVelocity } = calculateTrajectory(angle, velocity, airResistance, windSpeed);
    plotTrajectory(trajectory, range, timeOfFlight, finalVelocity);
});

function calculateTrajectory(angle, velocity, airResistance, windSpeed, dt = 0.1) {
    const angleRad = angle * (Math.PI / 180);
    let vx = velocity * Math.cos(angleRad) + windSpeed;
    let vy = velocity * Math.sin(angleRad);
    const g = 9.81;
    let x = 0, y = 0;
    const trajectory = [{ x, y }];

    while (y >= 0) {
        const ax = -airResistance * vx;
        const ay = -g - airResistance * vy;
        vx += ax * dt;
        vy += ay * dt;
        x += vx * dt;
        y += vy * dt;
        if (y >= 0) {
            trajectory.push({ x, y });
        }
    }

    const timeOfFlight = trajectory.length * dt;
    const finalVelocity = Math.sqrt(vx ** 2 + vy ** 2);
    return { trajectory, range: x, timeOfFlight, finalVelocity };
}

function plotTrajectory(trajectory, range, timeOfFlight, finalVelocity) {
    const xValues = trajectory.map(point => point.x);
    const yValues = trajectory.map(point => point.y);

    const trace = {
        x: xValues,
        y: yValues,
        mode: 'lines',
        type: 'scatter',
        name: 'Projectile Path'
    };

    const layout = {
        title: 'Projectile Motion with Air Resistance and Wind Effect',
        xaxis: { title: 'Distance (m)' },
        yaxis: { title: 'Height (m)' },
        annotations: [{
            x: xValues[xValues.length - 1],
            y: yValues[yValues.length - 1],
            text: `Range: ${range.toFixed(2)} m<br>Time: ${timeOfFlight.toFixed(2)} s<br>Final Velocity: ${finalVelocity.toFixed(2)} m/s`,
            showarrow: true,
            arrowhead: 2,
            ax: -20,
            ay: -30
        }]
    };

    Plotly.newPlot('plot', [trace], layout);
}