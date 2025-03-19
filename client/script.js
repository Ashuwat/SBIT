const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.3, 1000);
const renderer = new THREE.WebGLRenderer({ antialias: true });
renderer.setSize(window.innerWidth, window.innerHeight);
renderer.setClearColor(0xffffff, 1); // Set background to white

document.body.appendChild(renderer.domElement);

const controls = new THREE.OrbitControls(camera, renderer.domElement);
controls.enablePan = false; // Disable side-to-side movement
controls.minPolarAngle = Math.PI / 4; // Restrict downward view
controls.maxPolarAngle = (3 * Math.PI) / 4; // Restrict upward view

const nodeGroup = new THREE.Group();
scene.add(nodeGroup);

const gridSize = 1;
const divisions = 10;
const gridHelperXY = new THREE.GridHelper(gridSize, divisions, 0x000000, 0x000000);
const gridHelperYZ = new THREE.GridHelper(gridSize, divisions, 0x000000, 0x000000);
const gridHelperXZ = new THREE.GridHelper(gridSize, divisions, 0x000000, 0x000000);

gridHelperXY.rotation.x = Math.PI / 2;
gridHelperYZ.rotation.z = Math.PI / 2;

scene.add(gridHelperXY);
scene.add(gridHelperYZ);
scene.add(gridHelperXZ);    

const axesHelper = new THREE.AxesHelper(1);
scene.add(axesHelper);

camera.position.set(0.5, 0.5, 2);
camera.lookAt(0.5, 0.5, 0.5);

function updateNodes(newNodes) {
    nodeGroup.clear(); // Remove old nodes
    
    newNodes.forEach(node => {
        const sphereGeometry = new THREE.SphereGeometry(0.02, 16, 16);
        const sphereMaterial = new THREE.MeshBasicMaterial({ color: 0xff0000 });
        const sphere = new THREE.Mesh(sphereGeometry, sphereMaterial);
        sphere.position.set(node[0], node[1], node[2]);
        nodeGroup.add(sphere);
    });
}

async function fetchData() {
    try {
        const response = await fetch("http://localhost:3030/data");
        if (!response.ok) throw new Error("Network error");

        const data = await response.json();
        console.log("Received data:", data);

        updateNodes(data);
    } catch (error) {
        console.error("Error fetching data:", error);
    }
}

const fetchDataBtn = document.createElement("button");
fetchDataBtn.innerText = "Fetch Dsata";
fetchDataBtn.style.position = "absolute";
fetchDataBtn.style.top = "10px";
fetchDataBtn.style.left = "10px";
fetchDataBtn.addEventListener("click", fetchData);
document.body.appendChild(fetchDataBtn);

function animate() {
    requestAnimationFrame(animate);
    controls.update();
    renderer.render(scene, camera);
}
animate();

window.addEventListener('resize', () => {
    renderer.setSize(window.innerWidth, window.innerHeight);
    camera.aspect = window.innerWidth / window.innerHeight;
    camera.updateProjectionMatrix();
});
