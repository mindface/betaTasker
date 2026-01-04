"use client"
import React, { useEffect, useRef, useCallback } from 'react'
import { VRMSchema, VRMSpringBoneColliderMesh, VRMSpringBoneColliderGroup } from '../model/VRMSchema'
import { GLTFLoader, GLTF } from 'three/examples/jsm/loaders/GLTFLoader';
import ContentControl from './ContentControl'

import * as THREE from 'three'

const glb_url = "gltf/Xbot.glb"

export default function SectionViewer() {
  const canvas = useRef<HTMLCanvasElement >(null)
  const top = useRef(30)
  const left = useRef(0)
  const botton = useRef(0)
  const sendModel = useRef<any>(null)
  let mixer = useRef<any>()
  const v3A = useRef(new THREE.Vector3())

  function createColliderMesh(radius:number, offset: THREE.Vector3): VRMSpringBoneColliderMesh {
    const colliderMesh = new THREE.Mesh(new THREE.SphereBufferGeometry(radius,8,4),new THREE.MeshBasicMaterial({ visible: false }))
    colliderMesh.position.copy(offset)
    colliderMesh.name = 'vrmColliderSphere'
    colliderMesh.geometry.computeBoundingSphere()
    return colliderMesh;
  }

  function switcher (switcherNumber:number, switcherTriger:string) {
    switch (switcherTriger) {
      case 'top':
        top.current = top.current + switcherNumber
        break;
      case 'left':
        left.current = left.current + switcherNumber
        break;
      case 'right':
        left.current = left.current - switcherNumber
        break;
      case 'botton':
        botton.current = botton.current - switcherNumber
        break;
    }
  }

  const controlerAction = function(caseAction:string) :void {
    let counter = 0
    const tumer = setInterval(() => {
      counter += 0.01
      switcher(counter,caseAction)
      if(counter >= 0.1){
        clearInterval(tumer)
        counter = 0
      }
    },100)
  }

  const onThree = useCallback(() => {

    function importColliderMeshGroups(gltf:GLTF,smAnimation:VRMSchema.SecondaryAnimation): VRMSpringBoneColliderGroup[] {
      const vrmColliderGruops = smAnimation.colliderGroups
      if(vrmColliderGruops === undefined) return [];

      const colliderGroups: VRMSpringBoneColliderGroup[] = [];
      vrmColliderGruops.forEach(async (colliderGroup) => {
        if( colliderGroup.node === undefined || colliderGroup.colliders === undefined ) return;
        const bone = await gltf.parser.getDependency('node',colliderGroup.node)
        const colliders: VRMSpringBoneColliderMesh[] = []
        colliderGroup.colliders.forEach((collider) => {
          if(
            collider.offset === undefined ||
            collider.offset.x === undefined ||
            collider.offset.y === undefined ||
            collider.offset.z === undefined ||
            collider.radius === undefined
          ) return;

          const offset = v3A.current.set(
            collider.offset.x,
            collider.offset.y,
            -collider.offset.z,
          )
          const colliderMesh = createColliderMesh(collider.radius, offset)
          bone.add(colliderMesh)
          colliders.push(colliderMesh)
        })
        const colliderMeshGroup = {
          node: colliderGroup.node,
          colliders
        }
        colliderGroups.push(colliderMeshGroup)
      })
      return colliderGroups;
    }

    const width = window.innerWidth
    const height = window.innerHeight

    const scene = new THREE.Scene()
    // scene.fog = new THREE.Fog( 0xa0a0a0, 10, 50 )
    const camera = new THREE.PerspectiveCamera(75,width/height)
    const renderer = new THREE.WebGLRenderer({canvas:canvas.current as HTMLCanvasElement,antialias: true})
    renderer.setClearColor('#ffffff')
    renderer.setPixelRatio(window.devicePixelRatio)
    renderer.setSize(width,height)
    renderer.shadowMap.enabled = true
    renderer.shadowMap.type = THREE.PCFSoftShadowMap
    camera.position.set(0, 0, +500)

    // const ambientLight = new THREE.HemisphereLight(0x888888, 0x0000FF, 3.0);
    // scene.add(ambientLight);
    const dirLight = new THREE.DirectionalLight( 0xffffff );
    dirLight.position.set( 3, 10, 10 );
    dirLight.castShadow = true;
    dirLight.shadow.camera.top = 2;
    dirLight.shadow.camera.bottom = - 2;
    dirLight.shadow.camera.left = - 2;
    dirLight.shadow.camera.right = 2;
    dirLight.shadow.camera.near = 0.1;
    dirLight.shadow.camera.far = 40;
    scene.add(dirLight);

    const loader = new GLTFLoader();
    loader.load(glb_url, async function(gltf) {
        sendModel.current = gltf.scene
        const vrmExt: VRMSchema.VRM | undefined = gltf.parser.json.extensions?.VRM;
        
        sendModel.current.scale.set(100.0, 100.0, 100.0)
        sendModel.current.position.set(0,30,0)
        scene.add(sendModel.current)
        sendModel.current.traverse( function ( object:any ) {
          if ( object.isMesh ) object.castShadow = true;
        });

        const animations = gltf.animations;
        // numAnimations = animations.length;
        
        if(animations && animations.length){
          mixer.current = new THREE.AnimationMixer( sendModel.current );
          for (let index = 0; index < animations.length; index++) {
						let clip = animations[ index ];
						// const name = clip.name;

            // const settings = baseActions[ clip.name ] || additiveActions[ clip.name ];
            const action = mixer.current.clipAction( clip );
            // console.log(action)
            // setWeight( action, settings.weight );
            // action.play();


            // mixer.current.clipAction(animations[index]).play()
          }
        }

        const schemaSecondaryAnimation: VRMSchema.SecondaryAnimation | undefined = vrmExt?.secondaryAnimation
        console.log(gltf.parser.extensions?.VRM)
        if(schemaSecondaryAnimation){
          const colliderGrops = await importColliderMeshGroups(gltf,schemaSecondaryAnimation)
          console.log("colliderGrops")
          console.log(colliderGrops)
        }

        scene.add(sendModel.current)

      },(xhr) => {
        console.log( `${( xhr.loaded / xhr.total * 100 )}% loaded` );
      },
      ( error ) => {
        console.error( 'An error happened', error );
      },
    )

    const geo = new THREE.BoxGeometry(100,100,100)
    const mate = new THREE.MeshNormalMaterial()
    const box = new THREE.Mesh(geo,mate)
    scene.add(box)

    const clock  = new THREE.Clock();

    tick()

    function tick() {
      box.rotation.y += 0.01
      if(sendModel.current){
        sendModel.current.rotation.set(left.current,top.current,botton.current)
      }

      if(mixer.current) {
        mixer.current.update(clock.getDelta());
      }

      // for ( let i = 0; i !== numAnimations; ++ i ) {
      //   const action = allActions[ i ];
      //   const clip = action.getClip();
      //   const settings = baseActions[ clip.name ] || additiveActions[ clip.name ];
      //   settings.weight = action.getEffectiveWeight();
      // }

      renderer.render(scene,camera)
      requestAnimationFrame(tick)
    }

  },[]);

  function setWeight( action:any, weight:number ) {
    action.enabled = true;
    action.setEffectiveTimeScale( 1 );
    action.setEffectiveWeight( weight );

  }

  function executeCrossFade( startAction:any, endAction:any, duration:any ) {

    // Not only the start action, but also the end action must get a weight of 1 before fading
    // (concerning the start action this is already guaranteed in this place)

    if ( endAction ) {

      setWeight( endAction, 1 );
      endAction.time = 0;

      if ( startAction ) {

        // Crossfade with warping
        startAction.crossFadeTo( endAction, duration, true );
      } else {
        // Fade in
        endAction.fadeIn( duration );
      }

    } else {
      // Fade out
      startAction.fadeOut( duration );
    }
  }

  useEffect(() => {
    window.addEventListener('load',() => {
      onThree()
    })
  },[onThree])

  return(
    <div className="canvas-box">
      <ContentControl controlAction={controlerAction} />
      <canvas ref={canvas}></canvas>
    </div>
  )
}
