// Setup scene, camera, and renderer
const scene = new THREE.Scene();
scene.background = new THREE.Color(0x222222); // ✅ Background color for visibility

const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
camera.position.set(0, 0, 5);  
camera.lookAt(0, 0, 0);  // ✅ Ensures the camera looks at the center

const renderer = new THREE.WebGLRenderer({ antialias: true });
renderer.setSize(window.innerWidth, window.innerHeight);
document.body.appendChild(renderer.domElement);

console.log("Canvas added:", document.body.contains(renderer.domElement));

// Add OrbitControls for interactivity
const controls = new THREE.OrbitControls(camera, renderer.domElement);

// Define a list of nodes (3D positions)
const nodes = [
    { x: 1, y: 2, z: 3 },
    { x: -2, y: 1, z: 0 },
    { x: 0, y: -1, z: -3 },
    { x: 3, y: 3, z: 3 }
];

const geometry = new THREE.BufferGeometry();
const positions = new Float32Array(nodes.flatMap(node => [node.x, node.y, node.z]));
console.log("Positions array:", positions);  // ✅ Debug output

geometry.setAttribute('position', new THREE.BufferAttribute(positions, 3));

const material = new THREE.PointsMaterial({ color: 0xff0000, size: 0.2 });
const points = new THREE.Points(geometry, material);
scene.add(points);

// Render loop
function animate() {
    requestAnimationFrame(animate);
    controls.update();
    renderer.render(scene, camera);
}
animate();

// Handle window resize
window.addEventListener('resize', () => {
    camera.aspect = window.innerWidth / window.innerHeight;
    camera.updateProjectionMatrix();
    renderer.setSize(window.innerWidth, window.innerHeight);
});

// Reset camera position when button is clicked
document.getElementById("reset").addEventListener("click", () => {
    camera.position.set(0, 0, 5);
    controls.target.set(0, 0, 0);
    controls.update();
});
