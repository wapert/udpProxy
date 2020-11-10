(function(e){function t(t){for(var a,s,i=t[0],l=t[1],c=t[2],p=0,d=[];p<i.length;p++)s=i[p],r[s]&&d.push(r[s][0]),r[s]=0;for(a in l)Object.prototype.hasOwnProperty.call(l,a)&&(e[a]=l[a]);u&&u(t);while(d.length)d.shift()();return o.push.apply(o,c||[]),n()}function n(){for(var e,t=0;t<o.length;t++){for(var n=o[t],a=!0,i=1;i<n.length;i++){var l=n[i];0!==r[l]&&(a=!1)}a&&(o.splice(t--,1),e=s(s.s=n[0]))}return e}var a={},r={app:0},o=[];function s(t){if(a[t])return a[t].exports;var n=a[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,s),n.l=!0,n.exports}s.m=e,s.c=a,s.d=function(e,t,n){s.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},s.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},s.t=function(e,t){if(1&t&&(e=s(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(s.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var a in e)s.d(n,a,function(t){return e[t]}.bind(null,a));return n},s.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return s.d(t,"a",t),t},s.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},s.p="/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],l=i.push.bind(i);i.push=t,i=i.slice();for(var c=0;c<i.length;c++)t(i[c]);var u=l;o.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"56d7":function(e,t,n){"use strict";n.r(t);var a=n("2b0e"),r=n("9f7b"),o=n.n(r),s=n("4776"),i=n.n(s),l=(n("b15b"),n("f9e3"),n("2dd8"),n("a058"),function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"app"}},[n("sidebar-menu",{attrs:{menu:e.menu}}),n("b-container",[n("router-view")],1)],1)}),c=[],u={data:function(){return{menu:[{header:!0,title:"Link Abstraction",hiddenOnCollapse:!1},{href:"/",title:"Home",icon:"fa fa-user"},{href:"/cats",title:"cats",icon:"fa fa-chart-area"},{href:"/dogs",title:"dogs",icon:"fa fa-chart-area"}]}}},p=u,d=(n("5c0b"),n("2877")),m=Object(d["a"])(p,l,c,!1,null,null,null),f=m.exports,b=n("8c4f"),g=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"home-view-container"},[n("h1",[e._v(" SDWAN Traffic Management ")]),n("h5",[e._v("Total number of Traffic containers for use: "+e._s(e.animalsCount))]),n("h3",[n("font-awesome-icon",{attrs:{icon:"cat"}}),e._v("\n      "+e._s(e.getAllCats.length)+" +\n      "),n("font-awesome-icon",{attrs:{icon:"dog"}}),e._v("\n      "+e._s(e.getAllDogs.length)+"\n    ")],1),n("button",{staticClass:"btn btn-primary",on:{click:e.togglePetForm}},[e._v("Add New Container")]),e.showPetForm?n("b-form",{on:{submit:function(t){return t.preventDefault(),e.handleSubmit(t)}}},[n("b-form-group",{attrs:{id:"exampleInputGroup2",label:"Pet's Name:","label-for":"exampleInput2"}},[n("b-form-input",{attrs:{id:"exampleInput2",type:"text",required:"",placeholder:"Enter name"},model:{value:e.formData.name,callback:function(t){e.$set(e.formData,"name",t)},expression:"formData.name"}})],1),n("b-form-group",{attrs:{id:"exampleInputGroup3",label:"Species:","label-for":"exampleInput3"}},[n("b-form-select",{attrs:{id:"exampleInput3",options:["cats","dogs"],required:""},model:{value:e.formData.species,callback:function(t){e.$set(e.formData,"species",t)},expression:"formData.species"}})],1),n("b-form-group",{attrs:{id:"exampleInputGroup2",label:"Pet's Age:","label-for":"exampleInput2"}},[n("b-form-input",{attrs:{id:"exampleInput2",type:"number",required:"",placeholder:"Enter age"},model:{value:e.formData.age,callback:function(t){e.$set(e.formData,"age",t)},expression:"formData.age"}})],1),n("b-button",{attrs:{type:"submit",variant:"primary"}},[e._v("Submit")]),n("b-button",{attrs:{type:"reset",variant:"danger"}},[e._v("Reset")])],1):e._e()],1)},h=[],v=(n("7f7f"),n("cebc")),_=n("2f62"),y={name:"home",data:function(){return{showPetForm:!1,formData:{name:"",age:0,species:null}}},computed:Object(v["a"])({},Object(_["c"])(["animalsCount","getAllCats","getAllDogs"])),methods:Object(v["a"])({},Object(_["b"])(["addPet"]),{togglePetForm:function(){this.showPetForm=!this.showPetForm},handleSubmit:function(){var e=this.formData,t=e.species,n=e.age,a=e.name,r={species:t,pet:{name:a,age:n}};this.addPet(r),this.formData={name:"",age:0,species:null}}})},w=y,O=Object(d["a"])(w,g,h,!1,null,null,null),x=O.exports,j=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"home-view-container"},[n("PetTable",{attrs:{species:"cats",pets:e.cats}})],1)},P=[],S=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("h1",[e._v(" Traffic containers of "+e._s(e.species))]),n("b-table",{attrs:{striped:"",hover:"",items:e.pets},scopedSlots:e._u([{key:"name",fn:function(t){return[n("router-link",{attrs:{to:"/pets/"+e.species+"/"+t.index}},[e._v("\n        "+e._s(t.value)+"\n      ")])]}}])})],1)},D=[],k={props:{species:String,pets:Array}},C=k,$=Object(d["a"])(C,S,D,!1,null,null,null),A=$.exports,T={components:{PetTable:A},data:function(){return{}},computed:Object(v["a"])({},Object(_["d"])(["cats"]))},I=T,E=Object(d["a"])(I,j,P,!1,null,null,null),F=E.exports,M=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"home-view-container"},[n("PetTable",{attrs:{species:"dogs",pets:e.dogs}})],1)},H=[],q={components:{PetTable:A},data:function(){return{}},computed:Object(v["a"])({},Object(_["d"])(["dogs"]))},G=q,N=Object(d["a"])(G,M,H,!1,null,null,null),J=N.exports,K=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"home-view-container"},[n("h1",[e._v(e._s(e.animal.name)+" ("+e._s(e.$route.params.species)+")")]),n("p",[e._v("Age: "+e._s(e.animal.age)+" years old")]),n("p",[e._v("Breed: "+e._s(e.animal.breed))])])},R=[],B={data:function(){return{animal:{}}},computed:Object(v["a"])({},Object(_["d"])(["cats","dogs"])),mounted:function(){var e=this[this.$route.params.species][this.$route.params.id];this.animal=e}},L=B,W=Object(d["a"])(L,K,R,!1,null,null,null),Z=W.exports;a["a"].use(b["a"]);var z=new b["a"]({mode:"history",base:"/",routes:[{path:"/",name:"home",component:x},{path:"/cats",name:"cats",component:F},{path:"/dogs",name:"dogs",component:J},{path:"/pets/:species/:id",name:"pet",component:Z}]}),Q=n("75fc"),U=[{name:"Fish",breed:"tuxedo",species:"cat",gender:"male",age:20,color:"black/white",weight:13,location:"fourside",notes:"Sweet kitty. He loves getting his belly rubbed."},{name:"Henry",breed:"tabby",species:"cat",gender:"male",age:20,color:"orange/white",weight:17,location:"threed",notes:"Super friendly"},{name:"Roger",breed:"tabby",species:"cat",gender:"male",age:20,color:"gray",weight:15,location:"threed",notes:"Super friendly"},{name:"Kitkat",breed:"bombay",species:"cat",gender:"female",age:.9,color:"black",weight:9,location:"threed",notes:"Super friendly"}],V=[{name:"Sheeba",breed:"collie",gender:"female",age:7,color:"black/white",weight:34,location:"fourside",notes:"Pure breed. Trained for competitions."},{name:"Hillary",breed:"mut",gender:"female",age:17,color:"orange/white",weight:37,location:"threed",notes:"Super friendly"},{name:"Zeus",breed:"afghan hound",gender:"male",age:9,color:"gray",weight:68,location:"threed",notes:"Super friendly"},{name:"Katie",breed:"golden retriever",gender:"female",age:2,color:"black",weight:44,location:"threed",notes:"Super friendly"}],X={cats:U,dogs:V,pets:[].concat(Object(Q["a"])(U),Object(Q["a"])(V))},Y={appendPet:function(e,t){var n=t.species,a=t.pet;e[n].push(a)}},ee={addPet:function(e,t){var n=e.commit;n("appendPet",t)}},te={animalsCount:function(e){return e.cats.length+e.dogs.length},getAllCats:function(e){return e.cats},getAllDogs:function(e){return e.dogs}};a["a"].use(_["a"]);var ne=new _["a"].Store({state:X,mutations:Y,actions:ee,getters:te}),ae=n("ecee"),re=n("c074"),oe=n("ad3d");ae["c"].add(re["a"],re["b"]),a["a"].component("font-awesome-icon",oe["a"]),a["a"].use(o.a),a["a"].config.productionTip=!1,a["a"].use(i.a),new a["a"]({router:z,store:ne,render:function(e){return e(f)}}).$mount("#app")},"5c0b":function(e,t,n){"use strict";var a=n("5e27"),r=n.n(a);r.a},"5e27":function(e,t,n){},a058:function(e,t,n){}});
//# sourceMappingURL=app.d9aaa2a4.js.map